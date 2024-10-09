import tensorflow as tf
from tensorflow import keras

model = keras.Sequential([
  keras.layers.Dense(128, activation='relu'),
  keras.layers.Dense(64, activation='relu'),
  keras.layers.Dense(1, activation='sigmoid')
])

model.compile(optimizer='adam',
              loss='binary_crossentropy',
              metrics=['accuracy'])

# Train the model with your data
model.fit(train_data, train_labels, epochs=10)
