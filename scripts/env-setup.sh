# Colors for console output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}Starting local development environment...${NC}"

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}Docker is not running. Please start Docker first.${NC}"
    exit 1
fi

# Check if .env file exists, if not create from example
if [ ! -f .env ]; then
    if [ -f .env.example ]; then
        cp .env.example .env
        echo -e "${GREEN}Created .env file from .env.example${NC}"
    else
        echo -e "${RED}No .env or .env.example file found${NC}"
        exit 1
    fi
fi

# Start postgres
echo -e "${GREEN}Starting postgres...${NC}"
docker-compose up -d postgres

# Wait for postgres to be healthy
echo -e "${GREEN}Waiting for postgres to be ready...${NC}"
sleep 5

# Run database migrations
# echo -e "${GREEN}Running database migrations...${NC}"
# pnpm migration:run

# Start the application
# echo -e "${GREEN}Starting Go application...${NC}"
# gow run main.go