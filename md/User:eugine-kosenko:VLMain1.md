    main1 m n = main1p m n Nil Nil
       where
          main1p 0 0 = ScalarProduct
          main1p i 0 a = main1p (i-1) 0 (Cons i a)
          main1p 0 j a b = main1p 0 (j-1) a (Cons j b)
          main1p i j a b = main1p (i-1) (j-1) (Cons i a) (Cons j b)
