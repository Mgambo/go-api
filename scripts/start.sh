# Colors for console output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}Setting up swagger...${NC}"
swag init

echo -e "${GREEN}Start the app...${NC}"
gow run main.go