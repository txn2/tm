FROM scratch
ENV PATH=/bin

COPY tm /bin/

WORKDIR /

ENTRYPOINT ["/bin/tm"]