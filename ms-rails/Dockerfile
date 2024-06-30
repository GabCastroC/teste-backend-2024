FROM ruby:3.0.0-alpine3.13

RUN apk --update --upgrade add \
      bash\
      sqlite-dev\
      build-base\
      tzdata

WORKDIR /ms-rails

COPY Gemfile Gemfile.lock /ms-rails/

RUN bundle install