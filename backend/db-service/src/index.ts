import express from "express";
import cors from "cors";
import helmet from "helmet";
import "express-async-errors";

const app = express();

app.use(express.json());
app.use(cors());
app.use(helmet());

// example health route
app.get("/health", (_, res) => {
  res.send("OK");
});

const port = process.env.DB_SERVICE_PORT || 3000;

// Example health route
app.get("/health", (_, res) => {
  res.send({ status: "Service Healthy" });
});

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});