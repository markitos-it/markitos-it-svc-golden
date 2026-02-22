#!/bin/bash

#:[.'.]:>- ===================================================================================
#:[.'.]:>- Marco Antonio - markitos devsecops kulture
#:[.'.]:>- The Way of the Artisan
#:[.'.]:>- markitos.es.info@gmail.com
#:[.'.]:>- 🌍 https://github.com/orgs/markitos-it/repositories
#:[.'.]:>- 🌍 https://github.com/orgs/markitos-public/repositories
#:[.'.]:>- 📺 https://www.youtube.com/@markitos_devsecops
#:[.'.]:>- ===================================================================================

RED='\033[0;91m'
GREEN='\033[0;92m'
YELLOW='\033[0;93m'
BLUE='\033[0;94m'
CYAN='\033[0;96m'
MAGENTA='\033[0;95m'
WHITE='\033[0;97m'
NC='\033[0m'

HOST="localhost"
PORT="3000"
DELAY=2
SERVICE="goldens.GoldenService"
RESULTS=()

print_header() {
    echo -e "\n${WHITE}🚀 gRPC Service Test Suite${NC}\n"
}

print_test() {
    local title=$1
    local info=$2

    echo -e "${CYAN}──────────────────────────────────────────────────────────${NC}"
    echo -e "${BLUE}── ${title}${NC}"
    if [ -n "$info" ]; then
        echo -e "${WHITE}── INFO: ${info}${NC}"
    fi
    echo -e "${CYAN}──────────────────────────────────────────────────────────${NC}"
}

print_success() {
    echo -e "${GREEN}✓ Success${NC}: $1"
}

print_error() {
    echo -e "${RED}✗ Error${NC}: $1"
}

print_info() {
    echo -e "${YELLOW}ℹ Info${NC}: $1"
}

print_separator() {
    echo -e "${CYAN}────────────────────────────────────────────────────────${NC}"
}

wait_for_next_test() {
    echo -e "${MAGENTA}⏳ Esperando $DELAY segundos antes del siguiente test...${NC}"
    sleep $DELAY
    print_separator
    echo -e "\n\n"
}

record_result() {
    local name=$1
    local status=$2

    if [ "$status" -eq 0 ]; then
        RESULTS+=("${GREEN}[OK]${NC}:>- ${name}")
    else
        RESULTS+=("${RED}[KO]${NC}:>- ${name}")
    fi
}

#:[.'.]:>- TEST 0: Verificar conexión al servidor
check_server() {
    print_test "TEST 0: Verificando conexión al servidor"
    
    if grpcurl -plaintext ${HOST}:${PORT} list &>/dev/null; then
        print_success "Servidor gRPC disponible en ${HOST}:${PORT}"
        return 0
    else
        print_error "No se puede conectar al servidor gRPC en ${HOST}:${PORT}"
        echo -e "${YELLOW}Asegúrate de que el servicio está corriendo:${NC}"
        echo -e "${CYAN}  docker-compose up${NC}"
        return 1
    fi
}

#:[.'.]:>- TEST 1: Obtener todos los registros
test_get_all_goldens() {
    print_test "TEST 1: GetAllGoldens - Obtener todos los registros"
    
    echo -e "${CYAN}Comando:${NC}"
    echo -e "  ${WHITE}grpcurl -plaintext ${HOST}:${PORT} ${SERVICE}/GetAllGoldens${NC}\n"
    
    echo -e "${CYAN}Respuesta:${NC}"
    
    if grpcurl -plaintext ${HOST}:${PORT} ${SERVICE}/GetAllGoldens; then
        status=0
    else
        status=1
    fi

    print_separator
    if [ "$status" -eq 0 ]; then
        print_success "GetAllGoldens completado"
    else
        print_error "GetAllGoldens falló"
    fi
    return $status
}

#:[.'.]:>- TEST 2: Obtener registro por ID
test_get_golden_by_id() {
    local GOLDEN_ID=$1
    local INFO_LINE=$2
    
    print_test "TEST 2: GetGoldenById - Obtener registro por ID" "$INFO_LINE"
    
    echo -e "${CYAN}Comando:${NC}"
    echo -e "  ${WHITE}grpcurl -plaintext -d '{\"id\":\"${GOLDEN_ID}\"}' ${HOST}:${PORT} ${SERVICE}/GetGoldenById${NC}\n"
    echo -e "${CYAN}Parametros:${NC}"
    echo -e "  ${WHITE}id:${GOLDEN_ID}${NC}"
    
    echo -e "${CYAN}Respuesta:${NC}"
    
    if grpcurl -plaintext -d "{\"id\":\"${GOLDEN_ID}\"}" ${HOST}:${PORT} ${SERVICE}/GetGoldenById; then
        status=0
    else
        status=1
    fi

    print_separator
    if [ "$status" -eq 0 ]; then
        print_success "GetGoldenById completado"
    else
        print_error "GetGoldenById falló - Verifica que el ID existe"
    fi
    return $status
}

#:[.'.]:>- TEST 3: Listar servicios disponibles
test_list_services() {
    print_test "TEST 3: Listar servicios disponibles"
    
    echo -e "${CYAN}Comando:${NC}"
    echo -e "  ${WHITE}grpcurl -plaintext ${HOST}:${PORT} list${NC}\n"
    
    echo -e "${CYAN}Respuesta:${NC}"
    
    if grpcurl -plaintext ${HOST}:${PORT} list; then
        status=0
    else
        status=1
    fi

    print_separator
    if [ "$status" -eq 0 ]; then
        print_success "Listado de servicios completado"
    else
        print_error "No se pudo listar los servicios"
    fi
    return $status
}

main() {
    print_header
    
    print_info "Conectando a ${CYAN}${HOST}:${PORT}${NC}"
    print_separator
    
    if ! check_server; then
        exit 1
    fi
    
    wait_for_next_test
    
    test_get_all_goldens
    record_result "TEST 1: GetAllGoldens" $?
    
    wait_for_next_test
    
    FIRST_ID=$(grpcurl -plaintext ${HOST}:${PORT} ${SERVICE}/GetAllGoldens 2>/dev/null | grep '"id"' | head -1 | grep -o '"[^"]*"' | sed 's/"//g' | tail -1)
    
    if [ -z "$FIRST_ID" ]; then
        INFO_LINE="No se encontraron registros. Usando un ID de ejemplo."
        FIRST_ID="example-id-123"
    else
        INFO_LINE="Usando el primer ID encontrado: ${FIRST_ID}"
    fi
    
    test_get_golden_by_id "$FIRST_ID" "$INFO_LINE"
    record_result "TEST 2: GetGoldenById" $?
    
    wait_for_next_test
    
    test_list_services
    record_result "TEST 3: Listar servicios" $?

    print_separator
    echo -e "\n\n${WHITE}Resumen${NC}"
    for item in "${RESULTS[@]}"; do
        echo -e "$item"
    done
    echo -e "\n${GREEN}✓ Suite de pruebas completada${NC}\n"
}

if ! command -v grpcurl &> /dev/null; then
    echo -e "${RED}Error: grpcurl no está instalado${NC}"
    echo -e "${YELLOW}Instálalo con:${NC}"
    echo -e "  ${CYAN}go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest${NC}"
    exit 1
fi

main
