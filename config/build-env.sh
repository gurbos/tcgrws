#! /bin/bash

# build-env.sh populates the environment.conf file with data
# needed to configure the web service.

CONFIG_FILE=config.env

ENV_VARS=""


# Host public DNS name
PUBLIC_HOSTNAME=`/opt/aws/bin/ec2-metadata -p | cut -f2 -d:`
if [ "$PUBLIC_HOSTNAME" == "" ]; then
    echo "ERROR: Could not get public host name!"
    exit 1
fi
ENV_VARS="${ENV_VARS}PUBLIC_HOSTNAME=$PUBLIC_HOSTNAME\n"

# Database host DNS name
DB_HOST=`aws rds describe-db-instances --query "DBInstances[?DBInstanceIdentifier=='webservicedb'].Endpoint.Address" --output text`
if [ "$DB_HOST" == "" ]; then
    echo "ERROR: Could not get database host name!"
    exit 1
fi
ENV_VARS="${ENV_VARS}DB_HOST=$DB_HOST\n"

# Database port number
DB_PORT=`aws rds describe-db-instances --query "DBInstances[?DBInstanceIdentifier=='webservicedb'].Endpoint.Port" --output text`
if [ "$DB_PORT" == "" ]; then
    echo "ERROR: Could not get database port number!"
    exit 1
fi
ENV_VARS="${ENV_VARS}DB_PORT=$DB_PORT\n"

# Database username
if [ "$DB_USER" == "" ]; then
    echo "ERROR: Could not get database user name!"
    exit 1
fi
ENV_VARS="${ENV_VARS}DB_USER=$DB_USER\n"

# Database user password
if [ "$DB_PASS" == "" ]; then
    echo "ERROR: Could not get database user password!"
    exit 1
fi
ENV_VARS="${ENV_VARS}DB_PASS=$DB_PASS\n"

# Database name
if [ "$DB_NAME" == "" ]; then
    echo "ERROR: Could not get database name!"
    exit 1
fi
ENV_VARS="${ENV_VARS}DB_NAME=$DB_NAME\n"

# Static content path
if [ "$STATIC_CONTENT" == "" ]; then
    echo "ERROR: Could not get static content path!"
    exit 1
fi
ENV_VARS="${ENV_VARS}STATIC_CONTENT=$STATIC_CONTENT\n"
cd ..
echo -e $ENV_VARS > $CONFIG_FILE