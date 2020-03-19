ARG PYTHON_VERSION
FROM python:${PYTHON_VERSION} AS python-base

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONFAULTHANDLER 1

FROM python-base AS python-deps

WORKDIR /code

RUN apt-get update \
    && apt-get -y install freetds-dev \
    && apt-get -y install unixodbc-dev

COPY requirements.txt ./

RUN pip install -r requirements.txt -i https://pypi.douban.com/simple