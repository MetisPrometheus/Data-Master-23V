import warnings
import flwr as fl
import numpy as np

from sklearn.linear_model import LogisticRegression
from sklearn.metrics import log_loss

import utils
import syft as sy


