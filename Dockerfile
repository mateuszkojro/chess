#publicly available docker image "python" on docker hub will be pulled

FROM python

#creating directory app in container (linux machine)

RUN mkdir /home/app

#copying server.py from local directory to container's app folder

COPY server.py /home/app/server.py

#exposing port 3333 still need -p ip:3333:3333 to ip to be exposed

EXPOSE 3333

#runing a script

CMD python /home/app/server.py "" 3333

