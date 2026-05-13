import { Router } from "express";

const router = Router();

/**
 * @swagger
 * /api/health:
 *   get:
 *     summary: Returns OK status
 *     responses:
 *       200:
 *         description: OK message
 */
router.get("/health", (_req, res) => {
  res.status(200).send("ok");
});

export default router;
