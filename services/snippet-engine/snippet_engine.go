package snippetengine

import (
	"errors"
	"strings"
)

type CodeSnippet struct {
	ID             string   `json:"id"`
	Language       string   `json:"language"`
	Difficulty     string   `json:"difficulty"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Lines          []string `json:"lines"`
	ExpectedOutput string   `json:"expectedOutput,omitempty"`
	Tags           []string `json:"tags"`
}

var snippets = []CodeSnippet{
	// Python snippets
	{
		ID:          "py-basic-hello",
		Language:    "python",
		Difficulty:  "easy",
		Title:       "Hello World",
		Description: "Basic Python print statement",
		Lines: []string{
			"print(\"Hello, World!\")",
		},
		Tags: []string{"basic", "print"},
	},
	{
		ID:          "py-basic-variables",
		Language:    "python",
		Difficulty:  "easy",
		Title:       "Variable Assignment",
		Description: "Basic variable assignment and usage",
		Lines: []string{
			"name = \"Alice\"",
			"age = 25",
			"height = 5.6",
			"print(f\"Name: {name}, Age: {age}, Height: {height}\")",
		},
		Tags: []string{"variables", "f-strings"},
	},
	{
		ID:          "py-medium-function",
		Language:    "python",
		Difficulty:  "medium",
		Title:       "Function Definition",
		Description: "Function with parameters and return value",
		Lines: []string{
			"def calculate_area(length, width):",
			"    \"\"\"Calculate the area of a rectangle.\"\"\"",
			"    if length <= 0 or width <= 0:",
			"        raise ValueError(\"Dimensions must be positive\")",
			"    return length * width",
			"",
			"area = calculate_area(10, 5)",
			"print(f\"Area: {area}\")",
		},
		Tags: []string{"functions", "docstrings", "error-handling"},
	},
	{
		ID:          "py-hard-class",
		Language:    "python",
		Difficulty:  "hard",
		Title:       "Class Implementation",
		Description: "Class with inheritance and magic methods",
		Lines: []string{
			"class Animal:",
			"    def __init__(self, name, species):",
			"        self.name = name",
			"        self.species = species",
			"    ",
			"    def __str__(self):",
			"        return f\"{self.name} ({self.species})\"",
			"    ",
			"    def make_sound(self):",
			"        raise NotImplementedError",
			"",
			"class Dog(Animal):",
			"    def __init__(self, name, breed):",
			"        super().__init__(name, \"Canis lupus\")",
			"        self.breed = breed",
			"    ",
			"    def make_sound(self):",
			"        return \"Woof!\"",
		},
		Tags: []string{"classes", "inheritance", "magic-methods"},
	},

	// JavaScript snippets
	{
		ID:          "js-basic-hello",
		Language:    "javascript",
		Difficulty:  "easy",
		Title:       "Console Log",
		Description: "Basic JavaScript console output",
		Lines: []string{
			"console.log(\"Hello, World!\");",
		},
		Tags: []string{"basic", "console"},
	},
	{
		ID:          "js-basic-variables",
		Language:    "javascript",
		Difficulty:  "easy",
		Title:       "Variable Declaration",
		Description: "Modern JavaScript variable declaration",
		Lines: []string{
			"const name = \"Alice\";",
			"let age = 25;",
			"const height = 5.6;",
			"console.log(`Name: ${name}, Age: ${age}, Height: ${height}`);",
		},
		Tags: []string{"variables", "template-literals"},
	},
	{
		ID:          "js-medium-function",
		Language:    "javascript",
		Difficulty:  "medium",
		Title:       "Arrow Function",
		Description: "ES6 arrow function with error handling",
		Lines: []string{
			"const calculateArea = (length, width) => {",
			"    if (length <= 0 || width <= 0) {",
			"        throw new Error(\"Dimensions must be positive\");",
			"    }",
			"    return length * width;",
			"};",
			"",
			"try {",
			"    const area = calculateArea(10, 5);",
			"    console.log(`Area: ${area}`);",
			"} catch (error) {",
			"    console.error(error.message);",
			"}",
		},
		Tags: []string{"arrow-functions", "error-handling", "try-catch"},
	},
	{
		ID:          "js-hard-async",
		Language:    "javascript",
		Difficulty:  "hard",
		Title:       "Async/Await Pattern",
		Description: "Modern asynchronous JavaScript with error handling",
		Lines: []string{
			"class ApiClient {",
			"    constructor(baseURL) {",
			"        this.baseURL = baseURL;",
			"    }",
			"    ",
			"    async fetchUser(id) {",
			"        try {",
			"            const response = await fetch(`${this.baseURL}/users/${id}`);",
			"            if (!response.ok) {",
			"                throw new Error(`HTTP ${response.status}`);",
			"            }",
			"            return await response.json();",
			"        } catch (error) {",
			"            console.error(`Failed to fetch user ${id}:`, error);",
			"            throw error;",
			"        }",
			"    }",
			"}",
			"",
			"const client = new ApiClient(\"https://api.example.com\");",
			"client.fetchUser(123).then(user => console.log(user));",
		},
		Tags: []string{"async-await", "classes", "fetch", "error-handling"},
	},

	// Go snippets
	{
		ID:          "go-basic-hello",
		Language:    "go",
		Difficulty:  "easy",
		Title:       "Hello World",
		Description: "Basic Go program",
		Lines: []string{
			"package main",
			"",
			"import \"fmt\"",
			"",
			"func main() {",
			"    fmt.Println(\"Hello, World!\")",
			"}",
		},
		Tags: []string{"basic", "main", "fmt"},
	},
	{
		ID:          "go-basic-variables",
		Language:    "go",
		Difficulty:  "easy",
		Title:       "Variable Declaration",
		Description: "Go variable declaration and formatting",
		Lines: []string{
			"package main",
			"",
			"import \"fmt\"",
			"",
			"func main() {",
			"    var name string = \"Alice\"",
			"    age := 25",
			"    height := 5.6",
			"    fmt.Printf(\"Name: %s, Age: %d, Height: %.1f\\n\", name, age, height)",
			"}",
		},
		Tags: []string{"variables", "printf", "types"},
	},
	{
		ID:          "go-medium-function",
		Language:    "go",
		Difficulty:  "medium",
		Title:       "Function with Error",
		Description: "Go function with error handling",
		Lines: []string{
			"package main",
			"",
			"import (",
			"    \"errors\"",
			"    \"fmt\"",
			")",
			"",
			"func calculateArea(length, width float64) (float64, error) {",
			"    if length <= 0 || width <= 0 {",
			"        return 0, errors.New(\"dimensions must be positive\")",
			"    }",
			"    return length * width, nil",
			"}",
			"",
			"func main() {",
			"    area, err := calculateArea(10, 5)",
			"    if err != nil {",
			"        fmt.Printf(\"Error: %v\\n\", err)",
			"        return",
			"    }",
			"    fmt.Printf(\"Area: %.2f\\n\", area)",
			"}",
		},
		Tags: []string{"functions", "error-handling", "multiple-returns"},
	},
	{
		ID:          "go-hard-struct",
		Language:    "go",
		Difficulty:  "hard",
		Title:       "Struct with Methods",
		Description: "Go struct with methods and interfaces",
		Lines: []string{
			"package main",
			"",
			"import \"fmt\"",
			"",
			"type Shape interface {",
			"    Area() float64",
			"    Perimeter() float64",
			"}",
			"",
			"type Rectangle struct {",
			"    Length float64",
			"    Width  float64",
			"}",
			"",
			"func (r Rectangle) Area() float64 {",
			"    return r.Length * r.Width",
			"}",
			"",
			"func (r Rectangle) Perimeter() float64 {",
			"    return 2 * (r.Length + r.Width)",
			"}",
			"",
			"func printShapeInfo(s Shape) {",
			"    fmt.Printf(\"Area: %.2f, Perimeter: %.2f\\n\", s.Area(), s.Perimeter())",
			"}",
			"",
			"func main() {",
			"    rect := Rectangle{Length: 10, Width: 5}",
			"    printShapeInfo(rect)",
			"}",
		},
		Tags: []string{"structs", "methods", "interfaces", "polymorphism"},
	},
}

func GetSnippets(language, difficulty string, limit int) ([]CodeSnippet, error) {
	var filtered []CodeSnippet

	for _, snippet := range snippets {
		if language != "" && strings.ToLower(snippet.Language) != strings.ToLower(language) {
			continue
		}
		if difficulty != "" && strings.ToLower(snippet.Difficulty) != strings.ToLower(difficulty) {
			continue
		}
		filtered = append(filtered, snippet)
	}

	if limit > 0 && len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return filtered, nil
}

func GetSnippetByID(id string) (*CodeSnippet, error) {
	for _, snippet := range snippets {
		if snippet.ID == id {
			return &snippet, nil
		}
	}
	return nil, errors.New("snippet not found")
}

func GetDifficulties() []string {
	return []string{"easy", "medium", "hard"}
}

func GetLanguages() []string {
	return []string{"python", "javascript", "go"}
}