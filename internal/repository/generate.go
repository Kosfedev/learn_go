package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i QuestionRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i QuestionFormRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i QuestionFormUpdaterRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i QuestionSetRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i QuestionSubcategoryRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i DomainRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i CategoryRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i SubcategoryRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i SetRepository -o ./mocks -s "_minimock.go"
//go:generate minimock -i SetFormRepository -o ./mocks -s "_minimock.go"
