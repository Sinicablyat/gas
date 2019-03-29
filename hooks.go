package gas

// Hooks component lifecycle hooks
type Hooks struct {
	BeforeCreated HookWithControl // When parent already rendered (appended to DOM), but component Element don't yet (you can rerender childes)
	Created       Hook            // When component has been created in golang only (Element isn't available)

	Mounted Hook // When component has been mounted (Element is available)

	BeforeDestroy Hook // Before component destroy (Element is available)

	BeforeUpdate Hook // When component child don't updated
	Updated      Hook // After component child was updated
}

// Hook - lifecycle hook
type Hook func(*Component) error

// HookWithControl - lifecycle hook. Return true for rerender component childes
type HookWithControl func(this *Component) (rerender bool, err error)

// CallBeforeCreatedIfCan call component and it's childes BeforeCreated hook
func CallBeforeCreatedIfCan(i interface{}) error {
	if !IsComponent(i) {
		return nil
	}

	c := I2C(i)
	if c.Hooks.BeforeCreated != nil {
		rerender, err := c.Hooks.BeforeCreated(c)
		if err != nil {
			return err
		}

		if rerender {
			c.RChildes = RenderTree(c)
		}
	}

	for _, child := range c.RChildes {
		err := CallBeforeCreatedIfCan(child)
		if err != nil {
			return err
		}
	}

	return nil
}

// CallMountedIfCan call component and it's childes Mounted hook
func CallMountedIfCan(i interface{}) error {
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

		err := CallMountedIfCan(I2C(child))
		if err != nil {
			return err
		}
	}

	return nil
}

// CallBeforeDestroyIfCan call component and it's childes WillDestroy hook
func CallBeforeDestroyIfCan(i interface{}) error {
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

		err := CallBeforeDestroyIfCan(I2C(child))
		if err != nil {
			return err
		}
	}

	return nil
}

// CallUpdatedIfCan call component parent (true component) Updated hook
func CallUpdatedIfCan(i interface{}) error {
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

// CallBeforeUpdateIfCan call component parent (true component) BeforeUpdate
func CallBeforeUpdateIfCan(i interface{}) error {
	if !IsComponent(i) {
		return nil
	}

	// run BeforeUpdate hook for component parent(!)
	c := I2C(i).ParentComponent()

	if c.Hooks.BeforeUpdate != nil {
		err := c.Hooks.BeforeUpdate(c)
		if err != nil {
			return err
		}
	}

	return nil
}
