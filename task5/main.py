import matplotlib.pyplot as plt
import numpy as np
import tensorflow as tf
from keras.datasets import fashion_mnist

from tensorflow import keras

LIMIT = 200


def generation_dataset():
    return fashion_mnist.load_data()


def normalization_images(images):
    return images / 255.0


def create_model():
    model = tf.keras.Sequential([
        tf.keras.layers.Reshape((28, 28, 1), input_shape=(28, 28)),
        tf.keras.layers.Conv2D(32, (3, 3), activation='relu'),
        tf.keras.layers.MaxPooling2D((2, 2)),

        tf.keras.layers.Flatten(),
        # tf.keras.layers.Flatten(input_shape=(28, 28)),
        tf.keras.layers.Dense(128, activation='relu'),
        tf.keras.layers.Dense(1, activation='sigmoid')
    ])
    return model


def compile_model(model):
    model.compile(optimizer='adam',
                  loss=tf.keras.losses.BinaryCrossentropy(),
                  metrics=['accuracy'])

    return model


def training_model(model, train_images, train_labels):
    model.fit(train_images, train_labels, epochs=10)
    return model


def limit_images(items, limit=25):
    return items[:limit]



def main():
    print("Список фізичних пристроїв:", tf.config.list_physical_devices('GPU'))
    (train_images, train_labels), (test_images, test_labels) = generation_dataset()
    class_names = ['T-shirt/top', 'Trouser', 'Pullover', 'Dress', 'Coat',
                   'Sandal', 'Shirt', 'Sneaker', 'Bag', 'Ankle boot']

    # спрощуємо зображення, за рахунок нормалізації
    train_images = normalization_images(train_images)
    test_images = normalization_images(test_images)

    key_id = 3
    # train_images = limit_images(train_images, LIMIT)
    # train_labels = limit_images(train_labels, LIMIT)

    # train_labels_filter = np.isin(train_labels, [key_id, 1])
    # test_labels_filter = np.isin(test_labels, [key_id, 1])

    train_images_filtered = train_images
    train_labels_filtered = train_labels
    test_images_filtered = test_images
    test_labels_filtered = test_labels



    plt.figure()
    plt.imshow(train_images[0])
    plt.colorbar()
    plt.grid(False)
    plt.show()

    train_images_filtered = limit_images(train_images_filtered, LIMIT)
    train_labels_filtered = limit_images(train_labels_filtered, LIMIT)
    test_images_filtered = limit_images(test_images_filtered, 15)
    test_labels_filtered = limit_images(test_labels_filtered, 15)



    train_labels_filtered = np.where(train_labels_filtered == key_id, 1, 0)
    test_labels_filtered = np.where(test_labels_filtered == key_id, 1, 0)

    print(train_images_filtered.shape)
    print(len(train_labels_filtered))
    print(train_labels_filtered)
    print(test_images_filtered.shape)
    print(len(test_labels_filtered))

    # показати усі тестові дані
    # plt.figure(figsize=(10, 10))
    # for i in range(25):
    #     plt.subplot(5, 5, i + 1)
    #     plt.xticks([])
    #     plt.yticks([])
    #     plt.grid(False)
    #     plt.imshow(train_images[i], cmap=plt.cm.binary)
    #     plt.xlabel(class_names[train_labels[i]])
    # plt.show()

    model = create_model()
    model = compile_model(model)

    model = training_model(model, train_images_filtered, train_labels_filtered)

    test_loss, test_acc = model.evaluate(test_images_filtered, test_labels_filtered, verbose=2)

    print('\nTest accuracy:', test_acc)

    print("\n--- Вивід інформації ---")

    predictions = model.predict(test_images_filtered)

    # Налаштування для графіку
    plt.figure(figsize=(12, 6))

    for i in range(len(test_images_filtered)):
        ax = plt.subplot(1, len(test_images_filtered), i + 1)

        # Показуємо картинку
        plt.imshow(test_images_filtered[i], cmap=plt.cm.binary)
        plt.xticks([])
        plt.yticks([])
        plt.grid(False)

        probability = predictions[i][0]  # Ймовірність (число від 0.0 до 1.0)
        predicted_label = 1 if probability >= 0.5 else 0
        true_label = test_labels_filtered[i]

        if predicted_label == true_label:
            color = 'green'
        else:
            color = 'red'


        label_text = (f"True: {true_label}\n"
                      f"Pred: {predicted_label}\n"
                      f"Prob: {probability:.2f}")

        plt.xlabel(label_text, color=color)

    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    main()
