
;; ==========================================
;; Lab 11 - Task 2
;; Exception Handling and Error Management
;; ==========================================

;; Custom function that deliberately throws an exception
(defn validate-age [age]
  (if (< age 18)
    (throw (ex-info "Age must be 18 or above"
                   {:age age :error "Underage"}))
    (println "Age validation successful:" age)))


;; 1. Handling ArithmeticException
(println "\n========== ARITHMETIC EXCEPTION ==========")

(try
  (println "Result:" (/ 10 0))

  (catch ArithmeticException e
    (println "Caught ArithmeticException:" (ex-message e))
    (println "Recovery: Returning default result 0")
    (println "Recovered result: 0")

    )

  (finally
    (println "Finally block: Arithmetic operation completed"))


;; 2. Handling NullPointerException
(println "\n========== NULL POINTER EXCEPTION ==========")

(try
  (let [text nil]
    (println "String length:" (.length ^String text)))

  (catch NullPointerException e
    (println "Caught NullPointerException:" (ex-message e))
    (println "Recovery: Using an empty string")
    (println "Recovered string length:" (.length "")))

  (finally
    (println "Finally block: Null check completed"))


;; 3. Handling a Custom Exception using ex-info
(println "\n========== CUSTOM EXCEPTION ==========")

(try
  (validate-age 15)

  (catch clojure.lang.ExceptionInfo e
    (println "Caught custom exception:" (ex-message e))
    (println "Exception data:" (ex-data e))
    (println "Age from exception data:" (:age (ex-data e)))
    (println "Error type:" (:error (ex-data e)))
    (println "Recovery: Allowing program to continue"))

  (finally
    (println "Finally block: Age validation completed"))


;; 4. Demonstrating successful validation
(println "\n========== SUCCESSFUL VALIDATION ==========")

(try
  (validate-age 21)

  (catch clojure.lang.ExceptionInfo e
    (println "Caught exception:" (ex-message e)))

  (finally
    (println "Finally block: Successful validation completed"))


(println "\nProgram completed without terminating unexpectedly.")
