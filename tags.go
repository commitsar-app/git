package git

func (g *Git) Tags() ([]*Tag, error) {
	tags, err := g.getTags()

	if err != nil {
		return nil, err
	}

	return tags, nil
}
