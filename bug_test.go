package bug

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"testing"

	"entgo.io/bug/ent/vehicle"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
)

func TestBugSQLite(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	test(t, client)
}

func TestBugMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()

	g1 := client.Garage.Create().SetName("Garage1").SaveX(ctx)
	g2 := client.Garage.Create().SetName("Garage2").SaveX(ctx)
	g3 := client.Garage.Create().SetName("Garage3").SaveX(ctx)
	g4 := client.Garage.Create().SetName("Garage4").SaveX(ctx)
	g5 := client.Garage.Create().SetName("Garage5").SaveX(ctx)
	g6 := client.Garage.Create().SetName("Garage6").SaveX(ctx)

	client.Plane.Create().SetName("Plane1").SetGarage(g1).SaveX(ctx)
	client.Plane.Create().SetName("Plane2").SetGarage(g2).SaveX(ctx)
	client.Plane.Create().SetName("Plane3").SetGarage(g3).SaveX(ctx)
	client.Car.Create().SetName("Car1").SetGarage(g4).SaveX(ctx)
	client.Car.Create().SetName("Car2").SetGarage(g5).SaveX(ctx)
	client.Car.Create().SetName("Car3").SetGarage(g6).SaveX(ctx)

	bq := client.Debug().Vehicle.Query()

	// subQuery (our virtual table that "union all" the car and plane type)
	subQuery := ent.Selector(ctx, bq.Clone().Modify(func(s *sql.Selector) {
		tb1 := sql.Table("cars").As("c")

		s.Select(tb1.C("id"), "'plane' AS `type`", tb1.C("name"), tb1.C("garage_id")).From(tb1)

		tb2 := sql.Table("planes").As("p")
		sel := sql.Select(tb2.C("id"), "'plane' AS `type`", tb2.C("name"), tb2.C("garage_id")).From(tb2)

		s.UnionAll(sel).As("vehicles")
	}))

	// ALL
	vehs := bq.Clone().Modify(func(s *sql.Selector) {
		s.Select(s.C(vehicle.FieldID), s.C(vehicle.FieldType), s.C(vehicle.FieldName), s.C(vehicle.FieldGarageID)).From(subQuery).As("g")
	}).AllX(ctx)
	for _, v := range vehs {
		fmt.Printf("%s:%s: %s\n", v.ID, v.Type, v.Name)
	}

	// 2022/02/07 18:36:46 driver.Query: query=SELECT DISTINCT `vehicles`.`id`, `vehicles`.`type`, `vehicles`.`name` FROM (SELECT `c`.`id`, 'plane' AS `type`, `c`.`name` FROM `cars` AS `c` UNION ALL SELECT `p`.`id`, 'plane' AS `type`, `p`.`name` FROM `planes` AS `p`) AS `vehicles` args=[]
	// 1:plane: Car1
	// 2:plane: Car2
	// 3:plane: Car3
	// 1:plane: Plane1
	// 2:plane: Plane2
	// 3:plane: Plane3

	// ALL with garages
	vehs = bq.Clone().Modify(func(s *sql.Selector) {
		s.Select(s.C(vehicle.FieldID), s.C(vehicle.FieldType), s.C(vehicle.FieldName), s.C(vehicle.FieldGarageID)).From(subQuery).As("g")
	}).WithGarage().AllX(ctx)
	for _, v := range vehs {
		fmt.Printf("%s:%s: %s Garage => %s\n", v.ID, v.Type, v.Name, v.Edges.Garage.Name)
	}

	// 2022/02/07 19:25:24 driver.Query: query=SELECT DISTINCT `vehicles`.`id`, `vehicles`.`type`, `vehicles`.`name`, `vehicles`.`garage_id` FROM (SELECT `c`.`id`, 'plane' AS `type`, `c`.`name`, `c`.`garage_id` FROM `cars` AS `c` UNION ALL SELECT `p`.`id`, 'plane' AS `type`, `p`.`name`, `p`.`garage_id` FROM `planes` AS `p`) AS `vehicles` args=[]
	// 2022/02/07 19:25:24 driver.Query: query=SELECT DISTINCT `garages`.`id`, `garages`.`name` FROM `garages` WHERE `garages`.`id` IN (?, ?, ?, ?, ?, ?) args=[4 5 6 1 2 3]
	// 1:plane: Car1 Garage => Garage4
	// 2:plane: Car2 Garage => Garage5
	// 3:plane: Car3 Garage => Garage6
	// 1:plane: Plane1 Garage => Garage1
	// 2:plane: Plane2 Garage => Garage2
	// 3:plane: Plane3 Garage => Garage3

	// WHERE
	vehs = bq.Clone().Modify(func(s *sql.Selector) {
		s.Select(s.C(vehicle.FieldID), s.C(vehicle.FieldType), s.C(vehicle.FieldName)).From(subQuery).As("g")
	}).Where(func(selector *sql.Selector) {
		selector.Where(sql.HasSuffix("name", "2"))
	}).AllX(ctx)
	for _, v := range vehs {
		fmt.Printf("%s:%s: %s\n", v.ID, v.Type, v.Name)
	}

	// 2022/02/07 18:36:46 driver.Query: query=SELECT DISTINCT `vehicles`.`id`, `vehicles`.`type`, `vehicles`.`name` FROM (SELECT `c`.`id`, 'plane' AS `type`, `c`.`name` FROM `cars` AS `c` UNION ALL SELECT `p`.`id`, 'plane' AS `type`, `p`.`name` FROM `planes` AS `p`) AS `vehicles` WHERE `name` LIKE ? args=[%2]
	// 2:plane: Car2
	// 2:plane: Plane2

	// LIMIT OFFSET
	vehs = bq.Clone().Modify(func(s *sql.Selector) {
		s.Select(s.C(vehicle.FieldID), s.C(vehicle.FieldType), s.C(vehicle.FieldName)).From(subQuery).As("g")
	}).Limit(1).Offset(1).AllX(ctx)
	for _, v := range vehs {
		fmt.Printf("%s:%s: %s\n", v.ID, v.Type, v.Name)
	}

	// 2022/02/07 18:36:46 driver.Query: query=SELECT DISTINCT `vehicles`.`id`, `vehicles`.`type`, `vehicles`.`name` FROM (SELECT `c`.`id`, 'plane' AS `type`, `c`.`name` FROM `cars` AS `c` UNION ALL SELECT `p`.`id`, 'plane' AS `type`, `p`.`name` FROM `planes` AS `p`) AS `vehicles` LIMIT 1 OFFSET 1 args=[]
	// 2:plane: Car2
}
