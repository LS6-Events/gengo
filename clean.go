package gengo

func (s *Service) clean() error {
	mainPkg, err := s.GetMainPackageName()
	if err != nil {
		return err
	}

	for i := 0; i < len(s.Components); i++ {
		f := s.Components[i]

		s.handleSpecialType(&f)

		if f.Package == mainPkg {
			f.Package = "main"
		}

		for k, v := range f.StructFields {
			err := cleanField(k, v, mainPkg, f.StructFields)
			if err != nil {
				return err
			}
		}

		s.Components[i] = f
	}

	return nil
}

func cleanField(k string, f Field, mainPkg string, returnTypes map[string]Field) error {
	if f.Package == mainPkg {
		f.Package = "main"
	}

	for k, v := range f.StructFields {
		err := cleanField(k, v, mainPkg, f.StructFields)
		if err != nil {
			return err
		}
	}

	returnTypes[k] = f

	return nil
}
