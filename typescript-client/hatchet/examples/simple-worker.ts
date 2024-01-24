import Hatchet from '..';
import { Workflow } from '../workflow';

const hatchet = Hatchet.init();

const workflow: Workflow = {
  id: 'example',
  description: 'test',
  on: {
    event: 'user:create',
  },
  steps: [
    {
      name: 'step1',
      run: (input, ctx) => {
        console.log('executed step1!');
        return { step1: 'step1' };
      },
    },
    {
      name: 'step2',
      run: (input, ctx) => {
        console.log('executed step2!');
        return { step2: 'step2' };
      },
    },
  ],
};

const worker = hatchet.worker(workflow);
worker.start();
