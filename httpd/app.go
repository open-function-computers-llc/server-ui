package main

func runApp() error {
	s := server{}
	err := s.start()
	if err != nil {
		return err
	}
	return nil
}
