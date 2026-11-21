
;; ==========================================
;; Lab 11 - Task 1
;; Java Interoperability in Clojure
;; ==========================================

;; 1. Calling Java Static Methods
(println "\n========== JAVA STATIC METHODS ==========")

(println "Absolute value using Math/abs:" (Math/abs -25))
(println "Maximum value using Math/max:" (Math/max 10 20))
(println "Current time in milliseconds:" (System/currentTimeMillis))


;; 2. Creating Java Object Instances
(println "\n========== JAVA OBJECT CREATION ==========")

;; Creating an ArrayList object
(def java-list (new java.util.ArrayList))

;; Creating a HashMap object
(def java-map (new java.util.HashMap))

(println "Created Java ArrayList:" java-list)
(println "Created Java HashMap:" java-map)


;; 3. Invoking Instance Methods
(println "\n========== INSTANCE METHODS ==========")

;; ArrayList methods
(.add java-list "Apple")
(.add java-list "Banana")
(.add java-list "Cherry")

(println "ArrayList after add:" java-list)
(println "Element at index 1 using get:" (.get java-list 1))
(println "ArrayList size:" (.size java-list))

;; HashMap methods
(.put java-map "name" "Divya")
(.put java-map "course" "Clojure")
(.put java-map "language" "Java")

(println "HashMap after put:" java-map)
(println "Value of name using get:" (.get java-map "name"))
(println "Contains course key?" (.containsKey java-map "course"))


;; 4. Converting Between Clojure and Java Collections
(println "\n========== COLLECTION CONVERSION ==========")

;; Clojure Vector to Java ArrayList
(def clojure-vector [10 20 30 40])
(def converted-java-list (java.util.ArrayList. clojure-vector))

(println "Original Clojure vector:" clojure-vector)
(println "Converted Java ArrayList:" converted-java-list)

;; Java ArrayList back to Clojure Vector
(def converted-clojure-vector (vec converted-java-list))

(println "Java ArrayList converted back to Clojure vector:"
         converted-clojure-vector)

;; Clojure Map to Java HashMap
(def clojure-map {:a 1 :b 2 :c 3})
(def converted-java-map (java.util.HashMap. clojure-map))

(println "Original Clojure map:" clojure-map)
(println "Converted Java HashMap:" converted-java-map)

;; Java HashMap back to Clojure Map
(def converted-clojure-map (into {} converted-java-map))

(println "Java HashMap converted back to Clojure map:"
         converted-clojure-map)
