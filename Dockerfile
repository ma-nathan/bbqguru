FROM alpine

ADD files/start.sh /
ADD settings.ini /
ADD files/bbq /

CMD /start.sh
