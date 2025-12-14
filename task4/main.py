import numpy as np
import matplotlib.pyplot as plt
import cv2


def generate_photo():

    img_32 = np.zeros((32, 32))

    img_32[8:24, 8:24] = 255

    for i in range(32):
        img_32[i, i] = 255

    noise = np.random.randint(0, 50, (32, 32))
    img_32 = np.clip(img_32 + noise, 0, 255)

    plt.figure(figsize=(4, 4))
    plt.imshow(img_32, cmap='gray', vmin=0, vmax=255)
    plt.title("Generated 32x32 Image")
    plt.grid(False)
    plt.show()

    cv2.imwrite('image_32x32.png', img_32)


# Функція згортки (спрощено)
def convolve(image, kernel, stride=1):
    k_h, k_w = kernel.shape
    h, w = image.shape
    # Розрахунок розміру виходу
    out_h = (h - k_h) // stride + 1
    out_w = (w - k_w) // stride + 1
    output = np.zeros((out_h, out_w))

    for y in range(0, out_h):
        for x in range(0, out_w):
            # Вирізаємо шматок зображення
            region = image[y * stride: y * stride + k_h, x * stride: x * stride + k_w]
            # Множимо на ядро і сумуємо (формула згортки)
            output[y, x] = np.sum(region * kernel)

    # ReLU активація (опціонально, але згадується в теорії [cite: 323])
    return np.maximum(0, output)


def max_pooling(image, pool_size=2, stride=2):
    h, w = image.shape
    out_h = (h - pool_size) // stride + 1
    out_w = (w - pool_size) // stride + 1
    output = np.zeros((out_h, out_w))

    for y in range(out_h):
        for x in range(out_w):
            region = image[y * stride: y * stride + pool_size, x * stride: x * stride + pool_size]
            output[y, x] = np.max(region)
    return output

def task_2(img):
    filters = {
        "Identity": np.array([[0, 0, 0], [0, 1, 0], [0, 0, 0]]),
        "Edge Detection 1": np.array([[1, 0, -1], [0, 0, 0], [-1, 0, 1]]),
        "Edge Detection 2": np.array([[0, 1, 0], [1, -4, 1], [0, 1, 0]]),
        "Sharpen": np.array([[-1, -1, -1], [-1, 8, -1], [-1, -1, -1]]),
        "Custom": np.array([[-1, -2, -1], [0, 0, 0], [1, 2, 1]])  # Приклад фільтра Собеля
    }

    # 3. Визначаємо завдання з таблиці (Ядро -> Крок)
    tasks = [
        ("Identity", 1), ("Identity", 2), ("Identity", 3),
        ("Edge Detection 1", 1), ("Edge Detection 1", 2), ("Edge Detection 1", 3),
        ("Edge Detection 2", 1), ("Edge Detection 2", 2), ("Edge Detection 2", 3),
        ("Sharpen", 1), ("Sharpen", 2), ("Sharpen", 3),  # У таблиці для Sharpen різні варіанти, додамо всі
        ("Custom", 1)
    ]

    print(f"Вхідне зображення: {img.shape}")

    # 4. Проходимо по всіх завданнях
    for name, stride in tasks:
        kernel = filters[name]

        # Крок 1: Згортка
        conv_res = convolve(img, kernel, stride=stride)

        # Крок 2: Pooling (завжди stride=2 для пулінгу згідно завдання [cite: 390])
        pool_res = max_pooling(conv_res, stride=2)

        # Візуалізація
        plt.figure(figsize=(8, 4))

        # Зображення після згортки
        plt.subplot(1, 2, 1)
        plt.imshow(conv_res, cmap='gray')
        plt.title(f"{name} (Stride {stride})\nShape: {conv_res.shape}")
        plt.axis('off')

        # Зображення після Pooling
        plt.subplot(1, 2, 2)
        plt.imshow(pool_res, cmap='gray')
        plt.title(f"After MaxPool (Stride 2)\nShape: {pool_res.shape}")
        plt.axis('off')

        plt.suptitle(f"Filter: {name}, Stride: {stride}")
        plt.show()


def task_3_sequential(original_img):
    print("\n--- ЗАВДАННЯ 3: Послідовне застосування ---")

    # Вибираємо 2 фільтри
    # 1. Edge Detection (щоб знайти контури)
    kernel_1 = np.array([[1, 0, -1], [0, 0, 0], [-1, 0, 1]])

    # 2. Sharpen (щоб підсилити знайдені ознаки)
    kernel_2 = np.array([[-1, -1, -1], [-1, 8, -1], [-1, -1, -1]])

    # === ШАР 1 ===
    print("Шар 1: Edge Detection + MaxPool")
    # Згортка
    layer1_conv = convolve(original_img, kernel_1, stride=1)
    # Пулінг (зменшує 32x32 -> 15x15 приблизно)
    layer1_out = max_pooling(layer1_conv, stride=2)

    # === ШАР 2 ===
    # ВАЖЛИВО: Вхідні дані тут - це ВИХІД попереднього шару (layer1_out)
    print(f"Шар 2: Sharpen + MaxPool (Вхідний розмір: {layer1_out.shape})")

    # Згортка другого рівня
    layer2_conv = convolve(layer1_out, kernel_2, stride=1)

    # Пулінг другого рівня (зменшує 13x13 -> 6x6 приблизно)
    layer2_out = max_pooling(layer2_conv, stride=2)

    # === ВІЗУАЛІЗАЦІЯ ЕВОЛЮЦІЇ ===
    plt.figure(figsize=(12, 4))

    # 1. Оригінал
    plt.subplot(1, 3, 1)
    plt.imshow(original_img, cmap='gray')
    plt.title(f"Вхідне\n{original_img.shape}")
    plt.axis('off')

    # 2. Після першого шару
    plt.subplot(1, 3, 2)
    plt.imshow(layer1_out, cmap='gray')
    plt.title(f"Після Шару 1\n{layer1_out.shape}")
    plt.axis('off')

    # 3. Після другого шару (Кінцевий результат)
    plt.subplot(1, 3, 3)
    plt.imshow(layer2_out, cmap='gray')
    plt.title(f"Після Шару 2\n{layer2_out.shape}")
    plt.axis('off')

    plt.suptitle("Завдання 3: Послідовна обробка")
    plt.show()


def generate_rgb_pattern(size=32):
    """
    Генерує зображення 32x32, де канали мають різні фігури:
    Red: Вертикальні лінії
    Green: Горизонтальні лінії
    Blue: Суцільний квадрат по центру
    """
    img = np.zeros((size, size, 3), dtype=np.float32)

    # 1. Червоний канал (R) - Вертикальні смуги
    for x in range(0, size, 4):
        img[:, x:x + 2, 0] = 255  # Red channel

    # 2. Зелений канал (G) - Горизонтальні смуги
    for y in range(0, size, 4):
        img[y:y + 2, :, 1] = 255  # Green channel

    # 3. Синій канал (B) - Квадрат
    img[8:24, 8:24, 2] = 255  # Blue channel

    # Нормалізація для коректного відображення (0..1 для float або 0..255 для uint8)
    return img.astype(np.uint8)


def task_4_color_convolution():
    print("\n--- ЗАВДАННЯ 4: Кольорове зображення (3 канали) ---")

    # 1. Генеруємо RGB зображення
    img_rgb = generate_rgb_pattern()

    # 2. Визначаємо ФІЛЬТР, який складається з 3-х ЯДЕР
    # Ядро для Червоного: шукає вертикальні межі
    kernel_R = np.array([[1, 0, -1], [1, 0, -1], [1, 0, -1]])

    # Ядро для Зеленого: шукає горизонтальні межі
    kernel_G = np.array([[1, 1, 1], [0, 0, 0], [-1, -1, -1]])

    # Ядро для Синього: просто ігноруємо його (нулі), або Identity
    kernel_B = np.zeros((3, 3))

    # 3. Виконуємо згортку окремо для кожного каналу [cite: 142]
    # img_rgb[:, :, 0] - це Червоний канал
    # img_rgb[:, :, 1] - це Зелений канал
    # img_rgb[:, :, 2] - це Синій канал

    out_r = convolve(img_rgb[:, :, 0], kernel_R, stride=1)
    out_g = convolve(img_rgb[:, :, 1], kernel_G, stride=1)
    out_b = convolve(img_rgb[:, :, 2], kernel_B, stride=1)

    # 4. Сумуємо результати + Bias (Зміщення) [cite: 145, 148]
    bias = 0
    final_output = out_r + out_g + out_b + bias

    # Застосовуємо ReLU (відкидаємо від'ємні значення)
    final_output = np.maximum(0, final_output)

    # === ВІЗУАЛІЗАЦІЯ ===
    plt.figure(figsize=(10, 6))

    # Показуємо вхідне кольорове фото
    plt.subplot(2, 3, 1)
    plt.imshow(img_rgb)
    plt.title("Вхідне RGB (32x32)")
    plt.axis('off')

    # Показуємо результат обробки каналів
    plt.subplot(2, 3, 4)
    plt.imshow(out_r, cmap='gray')
    plt.title("Вихід ядра R\n(Вертикальні межі)")
    plt.axis('off')

    plt.subplot(2, 3, 5)
    plt.imshow(out_g, cmap='gray')
    plt.title("Вихід ядра G\n(Горизонтальні межі)")
    plt.axis('off')

    plt.subplot(2, 3, 6)
    plt.imshow(out_b, cmap='gray')
    plt.title("Вихід ядра B\n(Порожньо)")
    plt.axis('off')

    # Показуємо фінальний результат (сума)
    plt.subplot(1, 3, 3)  # Велике зображення праворуч
    plt.imshow(final_output, cmap='gray')
    plt.title("Результат\n(Sum: R+G+B)")
    plt.axis('off')

    plt.tight_layout()
    plt.show()

def main():
    # generate_photo()
    # # 2. Визначаємо фільтри з таблиці
    img = cv2.imread('image_32x32.png', cv2.IMREAD_GRAYSCALE)
    task_2(img)
    task_3_sequential(img)

    task_4_color_convolution()


if __name__ == "__main__":
    main()
