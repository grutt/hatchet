import sleep from '../sleep';
import HatchetPromise from './hatchet-promise';

describe('HatchetPromise', () => {
  it('should resolve the original promise if not canceled', async () => {
    const hatchetPromise = new HatchetPromise(
      new Promise((resolve) => {
        setTimeout(() => resolve('RESOLVED'), 500);
      })
    );
    const result = await hatchetPromise.promise;
    expect(result).toEqual('RESOLVED');
  });
  it('should resolve the cancel promise if canceled', async () => {
    const hatchetPromise = new HatchetPromise(
      new Promise((resolve) => {
        setTimeout(() => resolve('RESOLVED'), 500);
      })
    );

    const result = hatchetPromise.promise;
    setTimeout(() => {
      hatchetPromise.cancel();
    }, 100);

    try {
      await result;
      expect(true).toEqual(false); // this should not be reached
    } catch (e) {
      expect(e).toEqual(undefined);
    }
  });

  describe('timeoutToMs', () => {
    it('should convert ms to ms', () => {
      expect(HatchetPromise.timeoutToMs('100ms')).toEqual(100);
    });
    it('should convert s to ms', () => {
      expect(HatchetPromise.timeoutToMs('100s')).toEqual(100000);
    });
    it('should convert m to ms', () => {
      expect(HatchetPromise.timeoutToMs('100m')).toEqual(6000000);
    });
    it('should convert h to ms', () => {
      expect(HatchetPromise.timeoutToMs('100h')).toEqual(360000000);
    });
    it('should convert d to ms', () => {
      expect(HatchetPromise.timeoutToMs('100d')).toEqual(8640000000);
    });
  });

  describe('timeout', () => {
    it('should timeout after 100ms', async () => {
      const spy = jest.fn();

      const hatchetPromise = new HatchetPromise(
        new Promise((resolve) => {
          setTimeout(() => {
            spy();
            return resolve('RESOLVED');
          }, 500);
        }),
        { timeout: '100ms' }
      );

      const result = hatchetPromise.promise;

      try {
        await result;
        expect(true).toEqual(false); // this should not be reached
      } catch (e: any) {
        await sleep(600);
        expect(e.message).toEqual(`Step timed out after 100ms`);
        expect(spy).not.toHaveBeenCalled();
      }
    });

    it('should not timeout if resolved before timeout', async () => {
      const spy = jest.fn();

      const hatchetPromise = new HatchetPromise(
        new Promise((resolve) => {
          setTimeout(() => {
            spy();
            return resolve('RESOLVED');
          }, 50);
        }),
        { timeout: '100ms' }
      );

      try {
        const result = await hatchetPromise.promise;
        await sleep(200);
        expect(result).toEqual('RESOLVED');
        expect(spy).toHaveBeenCalled();
      } catch (e: any) {
        expect(true).toEqual(false); // this should not be reached
      }
    });

    it('should not timeout if rejected before timeout', async () => {
      const err = new Error('REJECTED');

      const hatchetPromise = new HatchetPromise(
        new Promise((resolve, reject) => {
          setTimeout(() => {
            return reject(err);
          }, 50);
        }),
        { timeout: '100ms' }
      );
      try {
        await hatchetPromise.promise;
        await sleep(200);
        expect(true).toEqual(false); // this should not be reached
      } catch (e: any) {
        expect(e).toEqual(err);
      }
    });

    it('should not timeout if canceled before timeout', async () => {
      const hatchetPromise = new HatchetPromise(
        new Promise((reject) => {
          setTimeout(() => {
            return reject('RESOLVED');
          }, 100);
        }),
        { timeout: '100ms' }
      );

      try {
        hatchetPromise.cancel('Canceled');
        await hatchetPromise.promise;
        expect(true).toEqual(false); // this should not be reached
      } catch (e: any) {
        await sleep(200);
        expect(e).toEqual('Canceled');
      }
    });
  });
});
