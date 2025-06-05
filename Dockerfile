
FROM ubuntu:22.04

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl wget gnupg2 lsb-release \
    git build-essential \
    supervisor \
    golang \
    nodejs npm \
    yarn

# Set up working directories
WORKDIR /app

# Install ClickHouse
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv E0C56BD4 && \
    echo "deb https://packages.clickhouse.com/deb stable main" | tee /etc/apt/sources.list.d/clickhouse.list && \
    apt-get update && apt-get install -y clickhouse-server clickhouse-client clickhouse-common-static

# Copy app code
COPY . .

# Install frontend dependencies
WORKDIR /app/frontend
RUN yarn install

# Build frontend if needed
# RUN yarn build

# Set Go backend build
WORKDIR /app/backend
RUN go mod tidy && go build -o signoz-server .

# Add Supervisor config
WORKDIR /app
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Expose ports
EXPOSE 8123 8080 3000

# Start all services
CMD ["/usr/bin/supervisord", "-n"]
