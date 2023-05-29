package mock

type Controller struct {
}

func (c *Controller) ReadCardNumber() (string, error) {
	return "000-000-000-000", nil
}
