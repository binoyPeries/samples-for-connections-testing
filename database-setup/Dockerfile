FROM mysql:8.3

# Create a non-root user with UID > 1024
RUN groupadd -g 10001 mygroup && \
    useradd -u 10001 -g mygroup -s /bin/bash -m myuser

# Create new writable directories
RUN mkdir -p /tmp/mysql /tmp/mysqld && \
    chown -R myuser:mygroup /tmp/mysql /tmp/mysqld

# Set environment variables
ENV MYSQL_DATADIR=/tmp/mysql
ENV MYSQL_ROOT_PASSWORD=rootpass
ENV MYSQL_DATABASE=mydb
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

# Copy init script (must be owned by mysql user for init)
COPY --chown=10001:10001 init.sql /docker-entrypoint-initdb.d/

# Switch to non-root user
USER 10001

# Start MySQL with custom datadir and socket path
# Start MySQL with all writable paths set to /tmp
CMD ["mysqld", "--datadir=/tmp/mysql", "--socket=/tmp/mysqld/mysqld.sock", "--pid-file=/tmp/mysqld/mysqld.pid", "--mysqlx-socket=/tmp/mysqld/mysqlx.sock"]

