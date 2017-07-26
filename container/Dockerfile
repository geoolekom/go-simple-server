FROM ubuntu:16.04
MAINTAINER Egor Komarov <e.komarov@corp.mail.ru>

RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ precise-pgdg main" > /etc/apt/sources.list.d/pgdg.list
RUN apt-get update && apt-get install -y sudo zip screen python-software-properties software-properties-common postgresql-9.6 postgresql-client-9.6 postgresql-contrib-9.6

USER postgres
RUN /etc/init.d/postgresql start && \
    psql --command "CREATE USER go WITH SUPERUSER PASSWORD 'go';" && \
    createdb -O go gosimpleserver && \
    echo "host all  all    127.0.0.1/0  md5" >> /etc/postgresql/9.6/main/pg_hba.conf && \
    echo "listen_addresses='*'" >> /etc/postgresql/9.6/main/postgresql.conf && \
    sed -i "s/^datestyle.*/datestyle = 'iso, dmy'/" /etc/postgresql/9.6/main/postgresql.conf

USER root
RUN mkdir -p /media/data

WORKDIR /usr/bin
COPY run.sh .
COPY bin/go-simple-server .
ADD sql ./sql

EXPOSE 80
CMD ["bash", "run.sh"]

