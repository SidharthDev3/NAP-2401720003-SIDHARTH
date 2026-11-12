;; List
;; (def my-list '(10 20 30))
;; 
;; (println "Original List:" my-list)
;; (println "First Element:" (first my-list))
;; (println "Rest of List:" (rest my-list))
;; 
;; (def new-list (conj my-list 5))
;; 
;; (println "New List:" new-list)
;; (println "Original List:" my-list)

;; Vector
(def my-vector [10 20 30])

(println "Original Vector:" my-vector)
(println "Element at index 1:" (get my-vector 1))
(println "Count:" (count my-vector))

(def new-vector (conj my-vector 40))

(println "New Vector:" new-vector)
(println "Original Vector:" my-vector)