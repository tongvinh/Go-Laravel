package data

import (
    up "github.com/upper/db/v4"
    "time"
)
// test struct
type test struct {
    ID        int       `db:"id,omitempty"`
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *test) Table() string {
    return "tests"
}

// GetAll gets all records from the database, using upper
func (t *test) GetAll(condition up.Cond) ([]*test, error) {
    collection := upper.Collection(t.Table())
    var all []*test

    res := collection.Find(condition)
    err := res.All(&all)
    if err != nil {
        return nil, err
    }

    return all, err
}

// Get gets one record from the database, by id, using upper
func (t *test) Get(id int) (*test, error) {
    var one test
    collection := upper.Collection(t.Table())

    res := collection.Find(up.Cond{"id": id})
    err := res.One(&one)
    if err != nil {
        return nil, err
    }
    return &one, nil
}

// Update updates a record in the database, using upper
func (t *test) Update(m test) error {
    m.UpdatedAt = time.Now()
    collection := upper.Collection(t.Table())
    res := collection.Find(m.ID)
    err := res.Update(&m)
    if err != nil {
        return err
    }
    return nil
}

// Delete deletes a record from the database by id, using upper
func (t *test) Delete(id int) error {
    collection := upper.Collection(t.Table())
    res := collection.Find(id)
    err := res.Delete()
    if err != nil {
        return err
    }
    return nil
}

// Insert inserts a model into the database, using upper
func (t *test) Insert(m test) (int, error) {
    m.CreatedAt = time.Now()
    m.UpdatedAt = time.Now()
    collection := upper.Collection(t.Table())
    res, err := collection.Insert(m)
    if err != nil {
        return 0, err
    }

    id := getInsertID(res.ID())

    return id, nil
}

// Builder is an example of using upper's sql builder
func (t *test) Builder(id int) ([]*test, error) {
    collection := upper.Collection(t.Table())

    var result []*test

    err := collection.Session().
        SQL().
        SelectFrom(t.Table()).
        Where("id > ?", id).
        OrderBy("id").
        All(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

