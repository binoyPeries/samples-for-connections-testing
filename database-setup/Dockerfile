FROM mysql:8.3

# Create a writable data directory
RUN mkdir -p /tmp/mysql && chown -R mysql:mysql /tmp/mysql

# Copy init script
COPY init.sql /docker-entrypoint-initdb.d/

# Set env vars
ENV MYSQL_ROOT_PASSWORD=rootpass
ENV MYSQL_DATABASE=mydb
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

# Override data directory via env
ENV MYSQL_DATADIR=/tmp/mysql

# # Create a non-root user with UID > 1024
# RUN groupadd -g 10001 mygroup && \
#     useradd -u 10001 -g mygroup -s /bin/bash -m myuser

# # Create a new writable data directory
# RUN mkdir -p /tmp/mysql && chown -R myuser:mygroup /tmp/mysql

# # Set environment variable so MySQL knows where to store data
# ENV MYSQL_DATADIR=/tmp/mysql

# # Copy init script (will run only on first container start)
# COPY init.sql /docker-entrypoint-initdb.d/

# # Set required MySQL environment variables (adjust as needed)
# ENV MYSQL_ROOT_PASSWORD=rootpass
# ENV MYSQL_DATABASE=mydb
# ENV MYSQL_USER=myuser
# ENV MYSQL_PASSWORD=mypassword

# # Run as non-root user
# USER 10001

# # Start MySQL with custom data dir
# CMD ["mysqld", "--datadir=/tmp/mysql"]
