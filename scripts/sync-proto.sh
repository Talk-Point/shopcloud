#!/usr/bin/env bash

# Proto file sync script for private GitHub repositories
# This script downloads proto files from private repos using gh CLI

# Exit on error
set -e

# Configuration
PROTO_BASE_DIR="./proto/shopcloud"

# Repository configurations - add more as needed
# Format: service_name|repo|branch|proto_path
REPOS=(
    "sagehub|Talk-Point/sagehub|master|proto/shopcloud/sagehub/v1/sagehub.proto"
    # "consolehub|Talk-Point/cloud-console|master|proto/shopcloud/consolehub/v1/consolehub.proto"  # No proto file in repo
    "contenthub|Talk-Point/contenthub|master|proto/shopcloud/contenthub/v1/contenthub.proto"
    "eventhub|Talk-Point/eventhub|master|proto/shopcloud/eventhub/v1/eventhub.proto"
    "heartbeathub|Talk-Point/heartbeathub|master|proto/shopcloud/heartbeathub/v1/heartbeathub.proto"
    "imagehub|Talk-Point/imagehub|master|proto/shopcloud/imagehub/v1/imagehub.proto"
    "mailhub|Talk-Point/mailhub|master|proto/shopcloud/mailhub/v2/mailhub.proto"
    "metahub|Talk-Point/metahub|master|proto/shopcloud/metahub/v1/metahub.proto"
    "notificationhub|Talk-Point/notificationhub|master|proto/shopcloud/notificationhub/v1/notificationhub.proto"
    "supplierhub|Talk-Point/supplierhub|master|proto/shopcloud/supplierhub/v1/supplierhub.proto"
    "warehousebrainhub|Talk-Point/warehousebrainhub|master|proto/shopcloud/warehousebrainhub/v1/warehousebrainhub.proto"
)

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo -e "${RED}Error: gh CLI is not installed. Please install it first.${NC}"
    echo "Visit: https://cli.github.com/"
    exit 1
fi

# Check if authenticated
if ! gh auth status &> /dev/null; then
    echo -e "${RED}Error: Not authenticated with GitHub CLI.${NC}"
    echo "Please run: gh auth login"
    exit 1
fi

# Function to find repo info by service name
find_repo_info() {
    local service=$1
    for repo_info in "${REPOS[@]}"; do
        IFS='|' read -r s_name s_repo s_branch s_path <<< "$repo_info"
        if [ "$s_name" = "$service" ]; then
            echo "$repo_info"
            return 0
        fi
    done
    return 1
}

# Function to sync a single proto file
sync_proto() {
    local repo_info=$1
    IFS='|' read -r service repo branch proto_path <<< "$repo_info"
    
    echo -e "${YELLOW}Syncing ${service}...${NC}"
    
    # Extract destination path from proto_path
    local dest_dir=$(dirname "$PROTO_BASE_DIR/${proto_path#proto/shopcloud/}")
    local dest_file=$(basename "$proto_path")
    
    # Create directory if it doesn't exist
    mkdir -p "$dest_dir"
    
    # Download the proto file using gh CLI
    if gh api repos/${repo}/contents/${proto_path}?ref=${branch} \
        -H "Accept: application/vnd.github.raw" > "${dest_dir}/${dest_file}"; then
        echo -e "${GREEN}✓ Successfully synced ${service}${NC}"
    else
        echo -e "${RED}✗ Failed to sync ${service}${NC}"
        return 1
    fi
}

# Main execution
main() {
    echo "Starting proto file synchronization..."
    echo "=================================="
    
    local failed=0
    local success=0
    
    # Process specific service if provided as argument
    if [ $# -eq 1 ]; then
        repo_info=$(find_repo_info "$1")
        if [ $? -eq 0 ]; then
            if sync_proto "$repo_info"; then
                ((success++))
            else
                ((failed++))
            fi
        else
            echo -e "${RED}Error: Unknown service '$1'${NC}"
            echo "Available services:"
            for repo_info in "${REPOS[@]}"; do
                IFS='|' read -r s_name s_repo s_branch s_path <<< "$repo_info"
                echo "  - $s_name"
            done
            exit 1
        fi
    else
        # Sync all services
        for repo_info in "${REPOS[@]}"; do
            if sync_proto "$repo_info"; then
                ((success++))
            else
                ((failed++))
            fi
        done
    fi
    
    echo "=================================="
    echo -e "Summary: ${GREEN}${success} succeeded${NC}, ${RED}${failed} failed${NC}"
    
    if [ $failed -gt 0 ]; then
        exit 1
    fi
}

# Show usage
usage() {
    echo "Usage: $0 [service]"
    echo ""
    echo "Sync proto files from private GitHub repositories."
    echo ""
    echo "Options:"
    echo "  service    Optional: specific service to sync (e.g., sagehub)"
    echo "             If not provided, all services will be synced"
    echo ""
    echo "Available services:"
    for repo_info in "${REPOS[@]}"; do
        IFS='|' read -r s_name s_repo s_branch s_path <<< "$repo_info"
        echo "  - $s_name"
    done
}

# Handle help flag
if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
    usage
    exit 0
fi

# Run main function
main "$@"