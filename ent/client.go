// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"entgo.io/bug/ent/migrate"

	"entgo.io/bug/ent/car"
	"entgo.io/bug/ent/garage"
	"entgo.io/bug/ent/plane"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Garage is the client for interacting with the Garage builders.
	Garage *GarageClient
	// Plane is the client for interacting with the Plane builders.
	Plane *PlaneClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Car = NewCarClient(c.config)
	c.Garage = NewGarageClient(c.config)
	c.Plane = NewPlaneClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Car:    NewCarClient(cfg),
		Garage: NewGarageClient(cfg),
		Plane:  NewPlaneClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Car:    NewCarClient(cfg),
		Garage: NewGarageClient(cfg),
		Plane:  NewPlaneClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Car.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Car.Use(hooks...)
	c.Garage.Use(hooks...)
	c.Plane.Use(hooks...)
}

// CarClient is a client for the Car schema.
type CarClient struct {
	config
}

// NewCarClient returns a client for the Car from the given config.
func NewCarClient(c config) *CarClient {
	return &CarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `car.Hooks(f(g(h())))`.
func (c *CarClient) Use(hooks ...Hook) {
	c.hooks.Car = append(c.hooks.Car, hooks...)
}

// Create returns a create builder for Car.
func (c *CarClient) Create() *CarCreate {
	mutation := newCarMutation(c.config, OpCreate)
	return &CarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Car entities.
func (c *CarClient) CreateBulk(builders ...*CarCreate) *CarCreateBulk {
	return &CarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Car.
func (c *CarClient) Update() *CarUpdate {
	mutation := newCarMutation(c.config, OpUpdate)
	return &CarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CarClient) UpdateOne(ca *Car) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCar(ca))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CarClient) UpdateOneID(id int) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCarID(id))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Car.
func (c *CarClient) Delete() *CarDelete {
	mutation := newCarMutation(c.config, OpDelete)
	return &CarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CarClient) DeleteOne(ca *Car) *CarDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CarClient) DeleteOneID(id int) *CarDeleteOne {
	builder := c.Delete().Where(car.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CarDeleteOne{builder}
}

// Query returns a query builder for Car.
func (c *CarClient) Query() *CarQuery {
	return &CarQuery{
		config: c.config,
	}
}

// Get returns a Car entity by its id.
func (c *CarClient) Get(ctx context.Context, id int) (*Car, error) {
	return c.Query().Where(car.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CarClient) GetX(ctx context.Context, id int) *Car {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CarClient) Hooks() []Hook {
	return c.hooks.Car
}

// GarageClient is a client for the Garage schema.
type GarageClient struct {
	config
}

// NewGarageClient returns a client for the Garage from the given config.
func NewGarageClient(c config) *GarageClient {
	return &GarageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `garage.Hooks(f(g(h())))`.
func (c *GarageClient) Use(hooks ...Hook) {
	c.hooks.Garage = append(c.hooks.Garage, hooks...)
}

// Create returns a create builder for Garage.
func (c *GarageClient) Create() *GarageCreate {
	mutation := newGarageMutation(c.config, OpCreate)
	return &GarageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Garage entities.
func (c *GarageClient) CreateBulk(builders ...*GarageCreate) *GarageCreateBulk {
	return &GarageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Garage.
func (c *GarageClient) Update() *GarageUpdate {
	mutation := newGarageMutation(c.config, OpUpdate)
	return &GarageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GarageClient) UpdateOne(ga *Garage) *GarageUpdateOne {
	mutation := newGarageMutation(c.config, OpUpdateOne, withGarage(ga))
	return &GarageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GarageClient) UpdateOneID(id string) *GarageUpdateOne {
	mutation := newGarageMutation(c.config, OpUpdateOne, withGarageID(id))
	return &GarageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Garage.
func (c *GarageClient) Delete() *GarageDelete {
	mutation := newGarageMutation(c.config, OpDelete)
	return &GarageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GarageClient) DeleteOne(ga *Garage) *GarageDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GarageClient) DeleteOneID(id string) *GarageDeleteOne {
	builder := c.Delete().Where(garage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GarageDeleteOne{builder}
}

// Query returns a query builder for Garage.
func (c *GarageClient) Query() *GarageQuery {
	return &GarageQuery{
		config: c.config,
	}
}

// Get returns a Garage entity by its id.
func (c *GarageClient) Get(ctx context.Context, id string) (*Garage, error) {
	return c.Query().Where(garage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GarageClient) GetX(ctx context.Context, id string) *Garage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GarageClient) Hooks() []Hook {
	return c.hooks.Garage
}

// PlaneClient is a client for the Plane schema.
type PlaneClient struct {
	config
}

// NewPlaneClient returns a client for the Plane from the given config.
func NewPlaneClient(c config) *PlaneClient {
	return &PlaneClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `plane.Hooks(f(g(h())))`.
func (c *PlaneClient) Use(hooks ...Hook) {
	c.hooks.Plane = append(c.hooks.Plane, hooks...)
}

// Create returns a create builder for Plane.
func (c *PlaneClient) Create() *PlaneCreate {
	mutation := newPlaneMutation(c.config, OpCreate)
	return &PlaneCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Plane entities.
func (c *PlaneClient) CreateBulk(builders ...*PlaneCreate) *PlaneCreateBulk {
	return &PlaneCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Plane.
func (c *PlaneClient) Update() *PlaneUpdate {
	mutation := newPlaneMutation(c.config, OpUpdate)
	return &PlaneUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlaneClient) UpdateOne(pl *Plane) *PlaneUpdateOne {
	mutation := newPlaneMutation(c.config, OpUpdateOne, withPlane(pl))
	return &PlaneUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlaneClient) UpdateOneID(id int) *PlaneUpdateOne {
	mutation := newPlaneMutation(c.config, OpUpdateOne, withPlaneID(id))
	return &PlaneUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Plane.
func (c *PlaneClient) Delete() *PlaneDelete {
	mutation := newPlaneMutation(c.config, OpDelete)
	return &PlaneDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlaneClient) DeleteOne(pl *Plane) *PlaneDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlaneClient) DeleteOneID(id int) *PlaneDeleteOne {
	builder := c.Delete().Where(plane.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlaneDeleteOne{builder}
}

// Query returns a query builder for Plane.
func (c *PlaneClient) Query() *PlaneQuery {
	return &PlaneQuery{
		config: c.config,
	}
}

// Get returns a Plane entity by its id.
func (c *PlaneClient) Get(ctx context.Context, id int) (*Plane, error) {
	return c.Query().Where(plane.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlaneClient) GetX(ctx context.Context, id int) *Plane {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PlaneClient) Hooks() []Hook {
	return c.hooks.Plane
}
