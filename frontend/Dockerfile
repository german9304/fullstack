FROM node:10.16

WORKDIR /fullstack-frontend/

COPY . /fullstack-frontend/

RUN npm -g install yarn 

RUN yarn install

CMD ["yarn", "run", "dev"]