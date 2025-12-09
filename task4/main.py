import numpy as np
import matplotlib.pyplot as plt
import cv2


def main():
    generate_photo()

def generate_photo():

    img_32 = np.zeros((32, 32))

    img_32[8:24, 8:24] = 255

    for i in range(32):
        img_32[i, i] = 255

    noise = np.random.randint(0, 50, (32, 32))
    img_32 = np.clip(img_32 + noise, 0, 255)

    # Відображення того, що вийшло
    plt.figure(figsize=(4, 4))
    plt.imshow(img_32, cmap='gray', vmin=0, vmax=255)
    plt.title("Generated 32x32 Image")
    plt.grid(False)
    plt.show()

    cv2.imwrite('my_test_image_32x32.png', img_32)


if __name__ == "__main__":
    main()
