# MultiSelector

Simple solution for working with multiple repositories (SQL, NoSQL etc) with one selection rules.

## Usage 

Let's say you want to test the function LoadUserAddresses :

```go

func main() {
    r := NewMongoRepository()

    entity := SomeEntity{}
    err := r.Fetch(selector.NewSelector([]string{"id > 1"}), &entity)
    if err != nil {
        continue
    }

}
```

In repository:

```go
func (r *MongoRepository) Fetch(
    selector multiselector.Selector, 
    entity interface{},
) error {
    ...
    // get collection here
	bson, err := selector.ToBson()
	if err != nil {
		return selector.ErrBadFormat
	}
	count, err := col.Find(bson).Count()
	if err != nil {
		return fmt.Errorf("error: %s", err.Error())
	}
	if count == 0 {
		return fmt.Errorf("not found")
	}
	if col.Find(bson).One(entity); err != nil {
		return fmt.Errorf("get upp: %w", err)
	}
	return nil
}

```