import matplotlib.pyplot as plt
import numpy as np
from keras.src.callbacks import EarlyStopping
from tensorflow.keras import models, layers, losses, optimizers, metrics
from tensorflow.keras.datasets import imdb
from tensorflow.keras.preprocessing import sequence


def create_basic_model(max_features, ):
    model = models.Sequential()
    model.add(layers.Embedding(max_features, 32))
    model.add(layers.LSTM(32))
    model.add(layers.Dense(1, activation='sigmoid'))
    return model


def create_modified_model(max_features, ):
    model = models.Sequential()
    model.add(layers.Embedding(max_features, 32))
    model.add(layers.Bidirectional(
        layers.GRU(32, return_sequences=True, dropout=0.3, recurrent_dropout=0.1)
    ))
    model.add(layers.GRU(32, dropout=0.3, recurrent_dropout=0.1))
    model.add(layers.Dense(1, activation='sigmoid'))
    return model


def compile_model(model):
    model.compile(
        optimizer=optimizers.RMSprop(),
        loss=losses.BinaryCrossentropy(),
        metrics=[metrics.BinaryAccuracy()]
    )

    return model


def plot_loss_comparison(history_base, history_mod):
    loss_base = history_base.history['loss']
    val_loss_base = history_base.history['val_loss']

    loss_mod = history_mod.history['loss']
    val_loss_mod = history_mod.history['val_loss']

    epochs_base = range(1, len(loss_base) + 1)
    epochs_mod = range(1, len(loss_mod) + 1)

    plt.figure(figsize=(14, 6))

    plt.subplot(1, 2, 1)
    plt.plot(epochs_base, loss_base, 'bo-', label='Train Loss (Базова)')
    plt.plot(epochs_base, val_loss_base, 'r*-', label='Val Loss (Базова)')
    plt.title('Базова модель: Втрати (Loss)')
    plt.xlabel('Епохи')
    plt.ylabel('Loss')
    plt.legend()
    plt.grid(True)

    plt.subplot(1, 2, 2)
    plt.plot(epochs_mod, loss_mod, 'bo-', label='Train Loss (Модифікована)')
    plt.plot(epochs_mod, val_loss_mod, 'g*-', label='Val Loss (Модифікована)')
    plt.title('Модифікована модель: Втрати (Loss)')
    plt.xlabel('Епохи')
    plt.ylabel('Loss')
    plt.legend()
    plt.grid(True)

    plt.tight_layout()
    plt.show()

    plt.figure(figsize=(10, 6))
    plt.plot(epochs_base, val_loss_base, 'r--', label='Валідація (Базова)')
    plt.plot(epochs_mod, val_loss_mod, 'g-', label='Валідація (Модифікована)')
    plt.title('Порівняння валідаційних втрат: Базова vs Модифікована')
    plt.xlabel('Епохи')
    plt.ylabel('Val Loss')
    plt.legend()
    plt.grid(True)
    plt.show()


def main():
    max_features = 10000
    max_len = 100

    print("Завантаження даних...")
    (x_train, y_train), (x_test, y_test) = imdb.load_data(num_words=max_features)

    print(f"Розмір тренувальної вибірки: {x_train.shape}")
    print(f"Розмір тестової вибірки: {x_test.shape}")

    # Падінг послідовностей (приведення до однакової довжини)
    x_train = sequence.pad_sequences(x_train, maxlen=max_len)
    x_test = sequence.pad_sequences(x_test, maxlen=max_len)
    print('x_train shape', x_train.shape)
    print('x_test shape', x_test.shape)

    y_train = np.asarray(y_train).astype('float32')
    y_test = np.asarray(y_test).astype('float32')
    print('train_labels shape', y_train.shape)
    print('test_labels shape', y_test.shape)

    print("\n--- Навчання базової моделі ---")
    model_base = create_basic_model(max_features)
    model_base = compile_model(model_base)

    x_val = x_train[:10000]
    y_val = y_train[:10000]
    partial_x_train = x_train[10000:]
    partial_y_train = y_train[10000:]
    print("Навчання...")
    history_base = model_base.fit(partial_x_train,
                                  partial_y_train,
                                  epochs=10,
                                  batch_size=16,
                                  validation_data=(x_val, y_val))

    print("Тестування моделі...")
    # Оцінка точності на тестових даних
    results_base = model_base.evaluate(x_test, y_test)

    print(f"\nТочність базової моделі: {results_base[1] * 100:.2f}%")

    print("\n--- Навчання модифікованої моделі ---")
    model_mod = create_modified_model(max_features)
    model_mod = compile_model(model_mod)

    print("Навчання...")
    my_callbacks = [
        EarlyStopping(monitor='val_loss', patience=3, restore_best_weights=True)
    ]
    history_mod = model_mod.fit(partial_x_train,
                                partial_y_train,
                                epochs=10,
                                batch_size=36,
                                validation_data=(x_val, y_val),
                                callbacks=my_callbacks
                                )

    print("Тестування моделі...")
    results_mod = model_mod.evaluate(x_test, y_test)
    print(f"\nТочність модифікованої моделі: {results_mod[1] * 100:.2f}%")

    plot_loss_comparison(history_base, history_mod)


if __name__ == '__main__':
    main()
