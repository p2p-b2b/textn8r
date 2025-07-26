package textn8r

import (
	"testing"
)

func TestUpperCaseNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO WORLD"},
		{"Hello World", "HELLO WORLD"},
		{"HELLO WORLD", "HELLO WORLD"},
	}

	for _, tt := range tests {
		result := UpperCaseNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("UpperCaseNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestLowerCaseNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO WORLD", "hello world"},
		{"Hello World", "hello world"},
		{"hello world", "hello world"},
	}

	for _, tt := range tests {
		result := LowerCaseNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("LowerCaseNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestTrimSpaceNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  hello world  ", "hello world"},
		{"hello world", "hello world"},
		{"  hello world", "hello world"},
		{"  hello       world    ", "hello       world"},
	}

	for _, tt := range tests {
		result := TrimSpaceNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("TrimSpaceNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveExtraSpaceNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello   world", "hello world"},
		{"hello world", "hello world"},
		{"  hello   world  ", "hello world"},
	}

	for _, tt := range tests {
		result := RemoveExtraSpaceNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveExtraSpaceNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveAllSpaceNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloworld"},
		{" hello world ", "helloworld"},
		{"hello   world", "helloworld"},
	}

	for _, tt := range tests {
		result := RemoveAllSpaceNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveAllSpaceNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveCarriageReturnNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello\rworld\n", "helloworld\n"},
		{"hello\r\nworld", "hello\nworld"},
		{"hello\nworld", "hello\nworld"},
	}

	for _, tt := range tests {
		result := RemoveCarriageReturnNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveCarriageReturnNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveNewLineNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello\nworld", "helloworld"},
		{"hello\n\nworld", "helloworld"},
		{"hello\n\nworld\n\n", "helloworld"},
		{"hello world", "hello world"},
		{" hello world  ", " hello world  "},
	}

	for _, tt := range tests {
		result := RemoveNewLineNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveNewLineNormalizer(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveTabNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello\tworld", "helloworld"},
		{"hello\t\tworld", "helloworld"},
		{"hello world", "hello world"},
		{" hello world\t\t", " hello world"},
	}

	for _, tt := range tests {
		result := RemoveTabNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveTabNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveNonAlphanumericNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello@world!", "helloworld"},
		{"hello world", "helloworld"},
		{"hello123", "hello123"},
	}

	for _, tt := range tests {
		result := RemoveNonAlphanumericNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveNonAlphanumericNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveTildesNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello~world", "helloworld"},
		{"hello world", "hello world"},
		{"hello~~world", "helloworld"},
	}

	for _, tt := range tests {
		result := RemoveTildesNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveTildesNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveDiacriticsNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"héllo wörld", "hllo wrld"},
		{"hello world", "hello world"},
		{"café", "caf"},
	}

	for _, tt := range tests {
		result := RemoveDiacriticsNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveDiacriticsNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveSpecialCharactersNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello@world!", "hello world "},
		{"hello world", "hello world"},
		{"hello#world$", "hello world "},
	}

	for _, tt := range tests {
		result := RemoveSpecialCharactersNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveSpecialCharactersNormalizer(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestReplaceSpecialCharactersNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello@world!", "-", "hello-world-"},
		{"hello world", "-", "hello world"},
		{"hello#world$", "_", "hello_world_"},
	}

	for _, tt := range tests {
		result := ReplaceSpecialCharactersNormalizer(tt.input, tt.replacement)
		if result != tt.expected {
			t.Errorf("ReplaceSpecialCharactersNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplaceAccentsNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"héllo wörld", "hello world"},
		{"café", "cafe"},
		{"mañana", "manana"},
		{"jalapeño", "jalapeno"},
		{"CAFÉ", "CAFE"},
		{"MAÑANA", "MANANA"},
	}

	for _, tt := range tests {
		result := ReplaceAccentsNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceAccentsNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestReplaceTildesNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"mañana", "manana"},
		{"niño", "nino"},
		{"jalapeño", "jalapeno"},
	}

	for _, tt := range tests {
		result := ReplaceTildesNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceTildesNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemovePunctuationNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello, world!", "hello world"},
		{"hello.world", "helloworld"},
		{"hello-world", "helloworld"},
	}

	for _, tt := range tests {
		result := RemovePunctuationNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemovePunctuationNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestRemoveDigitsNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello123", "hello"},
		{"123456", ""},
		{"hello world", "hello world"},
	}

	for _, tt := range tests {
		result := RemoveDigitsNormalizer(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveDigitsNormalizer(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestReplaceTabNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello\tworld", " ", "hello world"},
		{"hello\t\tworld", "-", "hello--world"},
		{"hello world", "_", "hello world"},
	}

	for _, tt := range tests {
		result := ReplaceTabNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceTabNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplaceCarriageReturnNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello\rworld\n", " ", "hello world\n"},
		{"hello\r\nworld", "-", "hello-\nworld"},
		{"hello\nworld", "_", "hello\nworld"},
	}

	for _, tt := range tests {
		result := ReplaceCarriageReturnNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceCarriageReturnNormalizer(%q, %q) = %q; want %q", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplaceNonAlphanumericNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello@world!", "-", "hello-world-"},
		{"hello world", "_", "hello_world"},
		{"hello#world$", "*", "hello*world*"},
	}

	for _, tt := range tests {
		result := ReplaceNonAlphanumericNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceNonAlphanumericNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplacePunctuationNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello, world!", " ", "hello  world "},
		{"hello.world", "-", "hello-world"},
		{"hello-world", "_", "hello_world"},
	}

	for _, tt := range tests {
		result := ReplacePunctuationNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplacePunctuationNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplaceDigitsNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello123", "-", "hello---"},
		{"123456", "*", "******"},
		{"hello world", "_", "hello world"},
	}

	for _, tt := range tests {
		result := ReplaceDigitsNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceDigitsNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplaceSpaceNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello world", "-", "hello-world"},
		{" hello world ", "_", "_hello_world_"},
		{"hello   world", "*", "hello***world"},
	}

	for _, tt := range tests {
		result := ReplaceSpaceNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceSpaceNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplaceDiacriticsNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"héllo wörld", "-", "h-llo w-rld"},
		{"café", "*", "caf*"},
		{"mañana", "_", "ma_ana"},
	}

	for _, tt := range tests {
		result := ReplaceDiacriticsNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceDiacriticsNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestReplaceNewLineNormalizer(t *testing.T) {
	tests := []struct {
		input       string
		replacement string
		expected    string
	}{
		{"hello\nworld", " ", "hello world"},
		{"hello\n\nworld", "-", "hello--world"},
		{"hello world", "_", "hello world"},
	}

	for _, tt := range tests {
		result := ReplaceNewLineNormalizer(tt.replacement).Apply(tt.input)
		if result != tt.expected {
			t.Errorf("ReplaceNewLineNormalizer(%s, %s) = %s; want %s", tt.input, tt.replacement, result, tt.expected)
		}
	}
}

func TestNormalizersApply(t *testing.T) {
	tests := []struct {
		input       string
		normalizers Normalizers
		expected    string
	}{
		{
			input: "  hello@world!  ",
			normalizers: Normalizers{
				TrimSpaceNormalizer,
				RemoveSpecialCharactersNormalizer,
				UpperCaseNormalizer,
			},
			expected: "HELLO WORLD ",
		},
		{
			input: "héllo wörld",
			normalizers: Normalizers{
				RemoveDiacriticsNormalizer,
				UpperCaseNormalizer,
			},
			expected: "HLLO WRLD",
		},
		{
			input: "hello\tworld",
			normalizers: Normalizers{
				RemoveTabNormalizer,
				UpperCaseNormalizer,
			},
			expected: "HELLOWORLD",
		},
	}

	for _, tt := range tests {
		result := tt.normalizers.Apply(tt.input)
		if result != tt.expected {
			t.Errorf("Normalizers.Apply(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestNormalizersApplyForSpanishTextWithNewLines(t *testing.T) {
	spanishText := `
¿Sabías que el número áureo, representado por la letra griega φ (phi), es aproximadamente 1,6180339887...?
Este número, presente en la naturaleza y el arte, ha fascinado a matemáticos y artistas
durante siglos. ¡Es increíble cómo algo tan abstracto puede manifestarse en la belleza de una flor o en
la proporción de una obra de arte! La secuencia de Fibonacci, donde cada número es la suma de los dos
anteriores (1, 1, 2, 3, 5, 8...), está íntimamente ligada al número áureo, y juntos forman una pareja matemática
que nos invita a reflexionar sobre la armonía del universo.

Ayer, mientras paseaba por el parque, observé a un grupo de niños jugando con un balón. ¡Qué energía y alegría
transmitían! Me pregunté: ¿cuántas veces habrán pateado ese balón? ¿Cuántos kilómetros habrán recorrido
en sus carreras? Me recordó a mi propia infancia, cuando el tiempo parecía detenerse mientras jugábamos al fútbol en la calle.
¡Ah, aquellos maravillosos años! Sin embargo, el tiempo no se detiene, y hoy me encuentro aquí, escribiendo estas líneas,
 recordando aquellos momentos y maravillándome ante la belleza de la vida.
`

	expectedText := `
SABIAS QUE EL NUMERO AUREO REPRESENTADO POR LA LETRA GRIEGA Φ PHI ES APROXIMADAMENTE 16180339887
ESTE NUMERO PRESENTE EN LA NATURALEZA Y EL ARTE HA FASCINADO A MATEMATICOS Y ARTISTAS
DURANTE SIGLOS ES INCREIBLE COMO ALGO TAN ABSTRACTO PUEDE MANIFESTARSE EN LA BELLEZA DE UNA FLOR O EN
LA PROPORCION DE UNA OBRA DE ARTE LA SECUENCIA DE FIBONACCI DONDE CADA NUMERO ES LA SUMA DE LOS DOS
ANTERIORES 1 1 2 3 5 8 ESTA INTIMAMENTE LIGADA AL NUMERO AUREO Y JUNTOS FORMAN UNA PAREJA MATEMATICA
QUE NOS INVITA A REFLEXIONAR SOBRE LA ARMONIA DEL UNIVERSO

AYER MIENTRAS PASEABA POR EL PARQUE OBSERVE A UN GRUPO DE NINOS JUGANDO CON UN BALON QUE ENERGIA Y ALEGRIA
TRANSMITIAN ME PREGUNTE CUANTAS VECES HABRAN PATEADO ESE BALON CUANTOS KILOMETROS HABRAN RECORRIDO
EN SUS CARRERAS ME RECORDO A MI PROPIA INFANCIA CUANDO EL TIEMPO PARECIA DETENERSE MIENTRAS JUGABAMOS AL FUTBOL EN LA CALLE
AH AQUELLOS MARAVILLOSOS ANOS SIN EMBARGO EL TIEMPO NO SE DETIENE Y HOY ME ENCUENTRO AQUI ESCRIBIENDO ESTAS LINEAS
 RECORDANDO AQUELLOS MOMENTOS Y MARAVILLANDOME ANTE LA BELLEZA DE LA VIDA
`

	result := Normalizers{
		ReplaceTildesNormalizer,
		ReplaceAccentsNormalizer,
		RemovePunctuationNormalizer,
		UpperCaseNormalizer,
	}.Apply(spanishText)

	if result != expectedText {
		t.Errorf("Normalizers.Apply()\n want:\n%q \n\ngot:\n%q", result, expectedText)
	}
}

func TestNormalizersApplyForSpanishTextWithoutNewLines(t *testing.T) {
	spanishText := `
¿Sabías que el número áureo, representado por la letra griega φ (phi), es aproximadamente 1,6180339887...?
Este número, presente en la naturaleza y el arte, ha fascinado a matemáticos y artistas
durante siglos.
`

	expectedText := ` SABIAS QUE EL NUMERO AUREO REPRESENTADO POR LA LETRA GRIEGA Φ PHI ES APROXIMADAMENTE 16180339887 ESTE NUMERO PRESENTE EN LA NATURALEZA Y EL ARTE HA FASCINADO A MATEMATICOS Y ARTISTAS DURANTE SIGLOS `

	result := Normalizers{
		ReplaceTildesNormalizer,
		ReplaceAccentsNormalizer,
		RemovePunctuationNormalizer,
		UpperCaseNormalizer,
		ReplaceNewLineNormalizer(" "),
	}.Apply(spanishText)

	if result != expectedText {
		t.Errorf("Normalizers.Apply()\n want:\n%q \n\ngot:\n%q", result, expectedText)
	}
}
