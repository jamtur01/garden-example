FROM ruby:2.5.3
MAINTAINER James Turnbull <james@lovedthanlost>

ENV service_path /mateservice
ENV PATH $PATH:$service_path

WORKDIR $service_path
COPY mateservice/ .

RUN bundle install

ENTRYPOINT ["ruby", "mate.rb"]

EXPOSE 8080
