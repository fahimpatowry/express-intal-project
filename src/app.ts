import express from "express";
import swaggerUi from "swagger-ui-express";
import swaggerSpec from "./config/swagger";

import healthRoute from "./routes/health.route";

const app = express();

app.use(express.json());

// Swagger setup
app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerSpec));

// routes
app.use("/", healthRoute);

export default app;