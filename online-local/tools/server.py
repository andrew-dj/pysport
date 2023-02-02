import logging
import subprocess
import sys

logging.root.addHandler(logging.StreamHandler())
logging.root.setLevel(logging.DEBUG)

if __name__ == '__main__':
    logging.info('Start Dev Server')
    try:
        go = subprocess.Popen(args=['go', 'run', './cmd/sportorg'], stdout=sys.stdout)
        node = subprocess.Popen(args=['npm', 'run', 'build'], cwd='./web', stdout=sys.stdout)
        go.wait()
        node.wait()
    except (KeyboardInterrupt, InterruptedError):
        pass
    logging.info('Stopping Dev Server')
