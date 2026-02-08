import express from "express";
import healthRoute from "./routes/health.route";

const app = express();

app.use(express.json());

// routes
app.use("/", healthRoute);

export default app;