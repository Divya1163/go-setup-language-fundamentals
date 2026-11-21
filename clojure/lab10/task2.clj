
;; ==========================================
;; Task 2: Sequence Transformations
;; map, filter, reduce, take, drop and ->>
;; ==========================================

;; COLLECTION 1: LIST
(println "\n========== COLLECTION 1: LIST ==========")

(def numbers-list '(1 2 3 4 5 6 7 8 9 10))

(println "Original list:" numbers-list)

;; 1. map - square each element
(println "Squares using map:"
         (map #(* % %) numbers-list))

;; 2. filter - retain only even numbers
(println "Even numbers using filter:"
         (filter even? numbers-list))

;; 3. reduce - calculate sum
(println "Sum using reduce:"
         (reduce + numbers-list))

;; 4. take - select first 5 elements
(println "First 5 elements using take:"
         (take 5 numbers-list))

;; 5. drop - skip first 5 elements
(println "After dropping first 5 elements:"
         (drop 5 numbers-list))

;; 6. Threading macro ->> - chain map and filter
(println "Threading ->> (square then filter even):"
         (->> numbers-list
              (map #(* % %))
              (filter even?)))


;; COLLECTION 2: VECTOR
(println "\n========== COLLECTION 2: VECTOR ==========")

(def numbers-vector [5 10 15 20 25 30])

(println "Original vector:" numbers-vector)

;; 1. map - add 10 to each element
(println "After map (add 10):"
         (map #(+ % 10) numbers-vector))

;; 2. filter - retain numbers greater than 15
(println "After filter (> 15):"
         (filter #(> % 15) numbers-vector))

;; 3. reduce - calculate maximum
(println "Maximum using reduce:"
         (reduce max numbers-vector))

;; 4. take - select first 3 elements
(println "First 3 elements using take:"
         (take 3 numbers-vector))

;; 5. drop - skip first 2 elements
(println "After dropping first 2 elements:"
         (drop 2 numbers-vector))

;; 6. Threading macro ->> - chain map and filter
(println "Threading ->> (add 10 then filter > 25):"
         (->> numbers-vector
              (map #(+ % 10))
              (filter #(> % 25))))
