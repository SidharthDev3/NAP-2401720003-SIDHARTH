;; List
(def my-list '(10 20 30))

(println "Original List:" my-list)
(println "First Element:" (first my-list))
(println "Rest of List:" (rest my-list))

(def new-list (conj my-list 5))

(println "New List:" new-list)
(println "Original List:" my-list)