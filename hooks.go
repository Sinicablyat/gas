package gas

// Hooks component lifecycle hooks
type Hooks struct {
	Created       Hook // When component has been created in golang only (Element isn't available)

	BeforeMounted Hook // When parent already rendered (appended to DOM), but component Element don't yet
	Mounted       Hook // When component has been mounted (Element is available)

	BeforeDestroy Hook // Before component destroy (Element is available)

	BeforeUpdate  Hook // When component child don't updated
	Updated       Hook // After component child was updated
}

// Hook - lifecycle hook
type Hook func(*Component) error

func RunMountedIfCan(i interface{}) error {
	if !IsComponent(i) {
		return nil
	}

	c := I2C(i)

	if c.Hooks.Mounted != nil {
		err := c.Hooks.Mounted(c)
		if err != nil {
			return err
		}
	}

	for _, child := range c.RChildes {
		if !IsComponent(child) {
			continue
		}

		err := RunMountedIfCan(I2C(child))
		if err != nil {
			return err
		}
	}

	return nil
}

func RunWillDestroyIfCan(i interface{}) error {
	if !IsComponent(i) {
		return nil
	}

	c := I2C(i)

	if c.Hooks.BeforeDestroy != nil {
		err := c.Hooks.BeforeDestroy(c)
		if err != nil {
			return err
		}
	}

	for _, child := range c.RChildes {
		if !IsComponent(child) {
			continue
		}

		err := RunWillDestroyIfCan(I2C(child))
		if err != nil {
			return err
		}
	}

	return nil
}

func RunUpdatedIfCan(i interface{}) error {
	if !IsComponent(i) {
		return nil
	}

	// run Updated hook for component parent(!)
	c := I2C(i).ParentComponent()

	if c.Hooks.Updated != nil {
		err := c.Hooks.Updated(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func RunBeforeUpdateIfCan(i interface{}) error {
	if !IsComponent(i) {
		return nil
	}

	// run BeforeUpdate hook for component parent(!)
	c := I2C(i).Parent

	if c.Hooks.BeforeUpdate != nil {
		err := c.Hooks.BeforeUpdate(c)
		if err != nil {
			return err
		}
	}

	return nil
}
