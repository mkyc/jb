build:
	cd ./tests/mock_apps && docker-compose build

compose:
	cd ./tests/mock_apps && docker-compose up --build
