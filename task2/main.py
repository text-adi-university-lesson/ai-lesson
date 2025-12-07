import numpy as np
import pandas as pd
global step
from graphviz import Digraph

# Розрахунок ентропії
def entropy(data):
    labels = data.iloc[:, -1]
    probabilities = labels.value_counts(normalize=True)
    return -sum(probabilities * np.log2(probabilities))


# Розрахунок інформаційного приросту
def information_gain(data, attribute):
    total_entropy = entropy(data)
    values = data[attribute].unique()
    subset_entropy = 0
    for value in values:
        subset = data[data[attribute] == value]
        weight = len(subset) / len(data)
        subset_entropy += weight * entropy(subset)
    return total_entropy - subset_entropy



# Вибір найкращого атрибута для розділення
def best_attribute(data, attributes):
    global step
    gains = [(attr, information_gain(data, attr)) for attr in attributes]
    step += 1
    print("Step:", step)
    for attr in gains:
        print(f"Attribute: {attr[0]}, Information Gain: {attr[1]:.4f}")
    return max(gains, key=lambda x: x[1])


# Побудова дерева рішень
def build_tree(data, attributes: pd.Index, target):
    labels = data[target]  # отримати позицію та вибір
    if len(labels.unique()) == 1:
        return labels.iloc[0]
    if not attributes.any():
        return labels.mode()[0]

    best_attr, _ = best_attribute(data, attributes)
    tree = {best_attr: {}}
    for value in data[best_attr].unique():
        subset = data[data[best_attr] == value]
        new_attrs = pd.Index(attr for attr in attributes if attr != best_attr)
        subtree = build_tree(subset, new_attrs, target)
        tree[best_attr][value] = subtree
    return tree


def predict(tree, example):
    if not isinstance(tree, dict):
        return tree  # якщо ми дійшли до листка → повертаємо результат
    root_attr = list(tree.keys())[0]  # корінь на поточному рівні
    subtree = tree[root_attr].get(example[root_attr])
    if subtree is None:
        # якщо немає значення в навчанні, можна взяти найбільш поширене
        return None
    return predict(subtree, example)

def add_nodes(graph, parent, data):
    if isinstance(data, dict):
        for key, value in data.items():
            graph.node(key)
            graph.edge(parent, key)
            add_nodes(graph, key, value)
    else:
        node_name = f"{parent}_{data}"
        graph.node(node_name, label=data, shape="box", color="lightblue")
        graph.edge(parent, node_name)

def main():
    global step
    step = 0
    data = pd.DataFrame(
        {'Price': ["Дешево", "Середньо", "Дешево", "Дорого", "Середньо", "Дешево", "Дорого", "Дешево",
                   "Дешево", "Середньо", "Середньо", "Дешево", "Дорого", "Середньо", "Дешево",
                   "Середньо", "Дорого", "Дешево", "Дорого", "Середньо"],
         'Brand': ['Dell', 'Apple', 'Asus', 'Apple', 'Dell', 'HP', 'Apple', 'Asus', 'Dell', 'HP',
                   'Apple', 'HP', 'Dell', 'Asus', 'HP', 'Apple', 'Dell', 'Asus', 'HP', 'Dell'],
         'RAM': [8, 16, 8, 16, 8, 8, 32, 8, 16, 16, 16, 8, 32, 16, 8, 16, 32, 8, 32, 16],
         'Weight': ['Легкий', 'Легкий', 'Легкий', 'Важкий', 'Легкий', 'Легкий', 'Важкий', 'Легкий',
                    'Легкий', 'Важкий', 'Легкий', 'Легкий', 'Важкий', 'Важкий', 'Легкий', 'Легкий',
                    'Важкий', 'Легкий', 'Важкий', 'Легкий'],
         'Type': ['Ігри', 'Професійне', 'Офіс', 'Професійне', 'Офіс', 'Ігри', 'Професійне', 'Офіс',
                  'Ігри', 'Професійне', 'Професійне', 'Офіс', 'Ігри', 'Професійне', 'Офіс',
                  'Професійне', 'Професійне', 'Ігри', 'Професійне', 'Офіс'],
         'Buy': ['Так', 'Так', 'Ні', 'Так', 'Так', 'Ні', 'Так', 'Ні', 'Так', 'Так', 'Так', 'Ні', 'Так',
                 'Так', 'Ні', 'Так', 'Так', 'Ні', 'Так', 'Так']})

    attributes = data.columns[:-1]
    target = 'Buy'

    train_data = data.iloc[:15]  # приклади для навчання
    test_data = data.iloc[15:]  # приклади для тесту

    tree = build_tree(train_data, attributes, target)
    print(tree)

    dot = Digraph()
    dot.node("Brand")
    add_nodes(dot, "Brand", tree["Brand"])
    dot.render("tree", format="png")

    predictions = []
    for i, row in test_data.iterrows():
        pred = predict(tree, row)
        predictions.append(pred)

    # Справжні значення
    actual = test_data['Buy'].tolist()

    # Вивід
    for i in range(len(predictions)):
        print(f"TestData {i + 1}: predicted={predictions[i]}, actual={actual[i]}")


if __name__ == "__main__":
    main()
