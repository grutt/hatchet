import HatchetError from '../errors/hatchet-error';

interface HatchetPromiseOptions {
  timeout: string;
}

type CancelFunction = (reason?: any) => void;

class HatchetPromise<T> {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  cancel: CancelFunction = (reason: any) => {};
  promise: Promise<T> | undefined;
  timeout: any | undefined;

  constructor(promise: Promise<T>, options?: HatchetPromiseOptions) {
    this.promise = new Promise((resolve, reject) => {
      this.cancel = (r: any = 'Canceled') => {
        this.clearTimeout();
        return reject(r);
      };

      Promise.resolve(promise)
        .then((value) => {
          resolve(value);
        })
        .catch((error) => {
          reject(error);
        })
        .finally(() => {
          this.clearTimeout();
        });
    });

    this.timeout = options?.timeout
      ? setTimeout(() => {
          this.cancel(new HatchetError(`Step timed out after ${options.timeout}`));
        }, HatchetPromise.timeoutToMs(options.timeout))
      : undefined;
  }

  async clearTimeout() {
    if (this.timeout) {
      clearTimeout(this.timeout);
    }
  }

  static timeoutToMs(timeout: string): number {
    const [value, unit] = timeout.split(/(?<=\d)(?=[a-zA-Z])/);
    switch (unit) {
      case 'ms':
        return parseInt(value, 10);
      case 's':
        return parseInt(value, 10) * 1000;
      case 'm':
        return parseInt(value, 10) * 1000 * 60;
      case 'h':
        return parseInt(value, 10) * 1000 * 60 * 60;
      case 'd':
        return parseInt(value, 10) * 1000 * 60 * 60 * 24;
      default:
        throw new HatchetError(`Invalid timeout unit: ${unit}`);
    }
  }
}

export default HatchetPromise;
