package main

import "testing"

func testOla(t *testing.T) {
	// primer test
	t.Run("diz ola pra diego", func(t *testing.T) {
		rest := greting("diego")
		output := "Ola, diego"

		if rest != output {
			t.Errorf("Resultado: %s\nResultado esperado %s", rest, output)
		}
	})
	// segundo
	t.Run("diz ola Mundo", func(t *testing.T) {
		rest := greting("")
		output := "Ola, Mundo"
		if rest != output {
			t.Errorf("Resultado: %s\nResultado esperado %s", rest, output)

		}
	})

}
