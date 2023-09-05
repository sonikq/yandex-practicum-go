/*

ID 89808402

-- ПРИНЦИП РАБОТЫ --

Начиная с корневого узла node, производится поиск узла с ключом key.

Если ключ текущего узла совпадает с key, то есть нашли узел для удаления:
	Если этот узел является листом, то он просто удаляется, возвращая nil.
	Если у узла есть только один дочерний узел (левый или правый), то этот дочерний узел становится на место удаляемого узла.
	Если у узла есть оба дочерних узла, то находится узел с минимальным значением в правом поддереве (самое левое значение в правом поддереве),
	это значение заменяет удаляемый узел, а затем узел с минимальным значением удаляется из правого поддерева.

Если ключ текущего узла меньше key, то поиск продолжается в правом поддереве.

Если ключ текущего узла больше key, то поиск продолжается в левом поддереве.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Удаление листового узла:
	Если узел, который нужно удалить, является листом, то его можно просто удалить из дерева, вернув nil.
	Это не нарушит свойства двоичного поискового дерева, так как нет необходимости пересортировывать другие узлы.

Удаление узла с одним дочерним узлом:
	Если у удаляемого узла есть только один дочерний узел,то этот дочерний узел можно поднять на место удаляемого узла.
	При этом порядок других узлов в дереве также сохраняется.

Удаление узла с двумя дочерними узлами:
	Если у удаляемого узла есть и левый, и правый дочерний узлы, то находится узел с минимальным значением в правом поддереве.
	Этот узел заменяет удаляемый узел, и затем узел с минимальным значением удаляется из правого поддерева.
	Таким образом, в левом поддереве не будет узлов со значениями больше, чем узел справа, и в правом поддереве не будет узлов со значениями меньше, чем узел слева.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Вставка, поиск и удаление узлов в сбалансированном двоичном поисковом дереве имеют среднюю временную сложность O(log n), где n - количество узлов в дереве.
Это происходит потому, что на каждом уровне дерева мы уменьшаем множество оставшихся узлов вдвое.

В худшем случае, когда дерево несбалансировано, операция удаления может иметь временную сложность O(n),
так как в этом случае каждая операция может затронуть все узлы на пути от корня до удаляемого узла.
Однако, при хорошо сбалансированных деревьях, этот худший случай редко происходит.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность зависит от рекурсивных вызовов и дополнительных переменных, которые создаются при каждом вызове функции remove и getAndDeleteMinVal.
В худшем случае, при сбалансированном дереве, глубина рекурсии будет O(log n), где n - количество узлов в дереве.

Таким образом, пространственная сложность составляет O(log n) для сбалансированного дерева.

*/

package B

type Node struct {
	value int
	left  *Node
	right *Node
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}
	if node.value == key {
		// если нужная нам нода это лист, то просто возвращаем nil
		if node.left == nil && node.right == nil {
			return nil
		}
		// если есть только правое или левое поддерево, то просто возвращаем эту поддерево
		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}
		// если есть и правое и левое поддерево, то ищем в правой самое маленькое значение
		// попутно удаляя его и ставя на место правого поддерева новое значение
		val, replaceNode := getAndDeleteMinVal(node.right)
		node.right = replaceNode
		node.value = val
		return node
	}
	if node.value > key {
		node.left = remove(node.left, key)
	} else {
		node.right = remove(node.right, key)
	}

	return node
}

func getAndDeleteMinVal(n *Node) (int, *Node) {
	if n.left == nil {
		if n.right == nil {
			return n.value, nil
		}
		return n.value, n.left
	}
	val, replaceNode := getAndDeleteMinVal(n.left)
	n.left = replaceNode
	return val, n
}
