FROM: centos:8

LABEL name="phrase counter"

RUN mkdir -p /phrase_counter/text

COPY parse-conc-arm64 /phrase_counter/
COPY samples /phrase_counter/samples

ENTRYPOINT ["/phrase_counter/parse-conc-arm64"]
CMD [""]