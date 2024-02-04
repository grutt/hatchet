import Hatchet from '../../src/sdk';
import { Workflow } from '../../src/workflow';

const hatchet = Hatchet.init();

const sleep = (ms: number) =>
  new Promise((resolve) => {
    setTimeout(resolve, ms);
  });

const workflow: Workflow = {
  id: 'concurrency-example',
  description: 'test',
  on: {
    event: 'concurrency:create',
  },
  concurrency: {
    name: 'user-concurrency',
    key: (ctx) => ctx.workflowInput().userId,
  },
  steps: [
    {
      name: 'step1',
      run: async (ctx) => {
        const { data } = ctx.workflowInput();
        console.log('starting step1 and waiting 5 seconds...', data);
        await sleep(5000);
        console.log('executed step1!');
        return { step1: `step1 results for ${data}!` };
      },
    },
    {
      name: 'step2',
      parents: ['step1'],
      run: (ctx) => {
        console.log('executed step2 after step1 returned ', ctx.stepOutput('step1'));
        return { step2: 'step2 results!' };
      },
    },
  ],
};

async function main() {
  const worker = await hatchet.worker('example-worker');
  await worker.registerWorkflow(workflow);
  worker.start();
}

main();
