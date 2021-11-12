import express from "express";
const app = express();

const usersRouter = require("./api/users");

app.use(express.json());
app.use("/users", usersRouter);

app.listen(3000);
