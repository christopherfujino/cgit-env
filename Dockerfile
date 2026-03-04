# bookworm, last updated 2026-02-28
FROM debian:13

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl --fail http://localhost/ || exit 1

# mailcap is for /etc/mime.types
RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install --no-install-recommends -y \
    cgit \
    lighttpd \
    mailcap \
  && apt-get clean

# Convenience dev deps
RUN apt-get install --no-install-recommends -y \
  vim \
  curl

COPY ./lighttpd.conf /etc/lighttpd/conf.d/cgit.conf
COPY ./cgitrc /etc/cgitrc

RUN echo 'include "conf.d/cgit.conf"' >> /etc/lighttpd/lighttpd.conf

COPY ./entrypoint.sh /entrypoint.sh

EXPOSE 80

ENTRYPOINT ["/entrypoint.sh"]
#CMD ["lighttpd", "-D", "-f", "/etc/lighttpd/lighttpd.conf"]
