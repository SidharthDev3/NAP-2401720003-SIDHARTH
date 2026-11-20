
(ns ExceptionHandling)


(defn check-age [age]
  (if (< age 18)
    (throw (ex-info "Age is less than 18" {:age age}))
    (println "Eligible")))


(try
  (/ 10 0)

  (catch ArithmeticException e
    (println "Arithmetic Error:" (.getMessage e)))

  (finally
    (println "Finally block executed")))



(try
  (.length nil)

  (catch NullPointerException e
    (println "Null Pointer Error"))

  (finally
    (println "Finally block executed")))



(try
  (check-age 15)

  (catch clojure.lang.ExceptionInfo e
    (println "Custom Error:" (ex-message e))
    (println "Error Data:" (ex-data e)))

  (finally
    (println "Finally block executed")))

(println "Program completed")
