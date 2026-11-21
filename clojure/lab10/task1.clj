
;; ==========================================
;; Demonstration of Core Immutable Data Structures
;; ==========================================

;; 1. LIST
(println "\n========== LIST ==========")

(def original-list '(10 20 30 40))

(println "Original list:" original-list)
(println "First element:" (first original-list))
(println "Rest of list:" (rest original-list))

(def new-list (conj original-list 5))

(println "After conj:" new-list)
(println "Original list remains unchanged:" original-list)


;; 2. VECTOR
(println "\n========== VECTOR ==========")

(def original-vector [10 20 30])

(println "Original vector:" original-vector)

(def new-vector (conj original-vector 40))

(println "After conj:" new-vector)
(println "Element at index 1:" (get original-vector 1))
(println "Count of elements:" (count original-vector))
(println "Original vector remains unchanged:" original-vector)


;; 3. MAP
(println "\n========== MAP ==========")

(def original-map {:name "Divya" :age 23 :course "Clojure"})

(println "Original map:" original-map)
(println "Value of name:" (get original-map :name))

(def updated-map (assoc original-map :city "Jaipur"))
(println "After assoc:" updated-map)

(def new-map (dissoc original-map :age))
(println "After dissoc:" new-map)

(println "Original map remains unchanged:" original-map)


;; 4. SET
(println "\n========== SET ==========")

(def original-set #{10 20 30})

(println "Original set:" original-set)
(println "Contains 20?" (contains? original-set 20))

(def new-set (conj original-set 40))
(println "After conj:" new-set)

(def removed-set (disj original-set 20))
(println "After disj:" removed-set)

(println "Original set remains unchanged:" original-set)
