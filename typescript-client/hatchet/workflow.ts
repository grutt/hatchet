import * as z from 'zod';

import { CreateStep, CreateStepSchema } from './step';

const CronConfigSchema = z.object({
  cron: z.string(), // TODO validate cron
  event: z.undefined(),
});

const EventConfigSchema = z.object({
  cron: z.undefined(),
  event: z.string(),
});

const OnConfigSchema = z.union([CronConfigSchema, EventConfigSchema]);

export const CreateWorkflowSchema = z.object({
  id: z.string(),
  description: z.string(),
  on: OnConfigSchema,
  steps: z.array(CreateStepSchema),
});

export interface Workflow extends z.infer<typeof CreateWorkflowSchema> {
  steps: CreateStep<any>[]; // TODO type this?
}
