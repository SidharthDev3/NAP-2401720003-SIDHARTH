
(ns JavaInteroperability
 (:import [java.util ArrayList HashMap]))


(println "Absolute value:" (Math/abs -10))
(println "Square root:" (Math/sqrt 25))
(println "Current time:" (System/currentTimeMillis))


(def list1 (ArrayList.))
(def map1 (HashMap.))


(.add list1 "Apple")
(.add list1 "Banana")
(.add list1 "Mango")

(.put map1 "Name" "Sidharth")
(.put map1 "Age" 20)

(println "Java ArrayList:" list1)
(println "Java HashMap:" map1)

(println "List size:" (.size list1))
(println "Map value:" (.get map1 "Name"))


(def clj-list [10 20 30 40])

(def java-list (ArrayList. clj-list))

(println "Clojure to Java:" java-list)


(def new-clj-list (vec java-list))

(println "Java to Clojure:" new-clj-list)
 